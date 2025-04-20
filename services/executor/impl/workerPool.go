package impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"sync"
)

type WorkerPool struct {
	taskChan      chan Instance   //待提交协程队列（带缓冲）
	semaphoreChan chan struct{}   //线程池信号量
	instanceMap   *sync.Map       //实例map
	ctx           context.Context //协程池ctx
	cancelFunc    context.CancelFunc
	l             logger.Logger
}

func NewWorkerPool(maxWorkerNum int) *WorkerPool {
	//父ctx来源于ServiceImpl的源context
	cancel, cancelFunc := context.WithCancel(context.Background())
	//子ctx：cancel用于控制协程池的生命周期
	return &WorkerPool{
		taskChan:      make(chan Instance, 500),          //任务队列默认上限500
		semaphoreChan: make(chan struct{}, maxWorkerNum), //信号量队列，用于控制线程池线程上限
		instanceMap:   &sync.Map{},
		ctx:           cancel,
		cancelFunc:    cancelFunc,
		l:             zap.L().Named("workerPool"),
	}
}

func (p *WorkerPool) submitInstance(task Instance) error {
	select {
	case <-p.ctx.Done():
		//	线程池关闭，无法继续写入任务
		//作为taskChan的写入方，关闭channel
		close(p.taskChan)
		return context.Canceled
	case p.taskChan <- task:
		task.recall(gen.InstanceStatus_WAITING_TO_SUBMIT, nil)
		return nil
	}
}

func (p *WorkerPool) ExecuteInstance(task Instance) {
	//	开始处理任务队列中的任务
	if task == nil {
		return
	}
	//	任务不为空,加入信号量，如果队列已满将阻塞，如果不满将可以继续加入
	select {
	case <-p.ctx.Done():
		close(p.semaphoreChan)
		p.l.Info("semaphore channel closed")
		return
	default:
		//正常来说，直接写入信号量
		p.semaphoreChan <- struct{}{}
		defer func() { <-p.semaphoreChan }() //结束后要把信号量拿出来
	}
	//通过map记录id->实例映
	instanceId, err := task.getId()
	if err != nil {
		return
	}
	p.instanceMap.Store(instanceId, task)
	defer func() {
		p.instanceMap.Delete(instanceId)
	}() //执行结束后要取出1个信号量并去掉map内的对应key
	//任务的执行看ServiceImpl的调用ctx,这里有点是如果线程池在这里被终止，这个任务会继续执行下去，因为
	//task的执行仅看ServiceImpl的调用ctx
	//任务回调，将运行中状态写入mysql
	task.recall(gen.InstanceStatus_EXECUTING, nil)
	err = task.execute()
	if err != nil {
		//	任务失败，回调失败
		if !errors.Is(err, context.Canceled) {
			task.recall(gen.InstanceStatus_FAIL, nil)
		}
		return
	}
	task.recall(gen.InstanceStatus_SUCCESS, nil)
}

func (p *WorkerPool) CancelInstance(ctx context.Context, instanceId string) error {
	//	取消某个instance比较简单，由于每个task本身有stop操作
	instance, ok := p.instanceMap.Load(instanceId)
	if !ok {
		//	case1 ExecuteInstance从未调用过、或者调用完成，instanceMap取不出该任务，已处理
		return fmt.Errorf("instance not found")
	}
	// case2  ExecuteInstance并发执行中
	//	分case1:由于并发已满，信号量无法写入，并case1处理，instanceMap取不出该任务
	//分case2:任务处于执行状态task.execute()，此时调用task.stop即可
	newInstance := instance.(Instance)
	_ = newInstance.stop() //传输信号一般理论不会失败
	p.instanceMap.Delete(instanceId)
	//回调取消
	newInstance.recall(gen.InstanceStatus_CANCEL, nil)
	return nil
}

func (p *WorkerPool) manageInstances() {
	for {
		select {
		case <-p.ctx.Done():
			//线程池退出
			p.l.Infof("workerPool exit")
			return

		case task, ok := <-p.taskChan:
			if !ok {
				p.l.Infof("task channel closed")
				return
			}
			//此时成功取得任务
			go p.ExecuteInstance(task)
		}
	}

}

func (p *WorkerPool) Start() {
	go p.manageInstances()
	p.l.Infof("workerPool start")
}

func (p *WorkerPool) Stop() {
	p.cancelFunc() //停止线程池后，线程池将停止接收处理任务
	//停止协程队列和信号量
	//close(p.taskChan)
	//close(p.semaphoreChan)
	//处理掉当前map里仍然存在的instance id，把仍然存在的instance 取消掉
	background := context.Background()
	p.instanceMap.Range(func(key, _ interface{}) bool {
		//停止
		_ = p.CancelInstance(background, key.(string)) //由于是线程池本身的stop操作，p.context已经被done了，这里手动取消instance则随便传一个
		return true
	})
	p.l.Infof("workerPool stop")
}
