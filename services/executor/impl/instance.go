package impl

//主要处理任务执行细节

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
)

type Instance interface {
	execute() error //执行任务
	stop() error
	getId() (string, error) //终止当前任务
}

type PythonInstance struct {
	*gen.TaskBase
	Id                string
	ctx               context.Context
	cancelFunc        context.CancelFunc
	pythonInterpreter string
}

func (p *PythonInstance) execute() error {
	err := kits.RunPythonScript(p.ctx, p.TaskBase.ScriptPath, "python_job_"+p.Id, "/logs", p.pythonInterpreter)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (p *PythonInstance) stop() error {
	if p.cancelFunc == nil || p.ctx == nil {
		return fmt.Errorf("cancel function or context is not exists")
	}
	//正常生成cancel func，开始发送取消信号
	p.cancelFunc()
	//检查是否获取到context结束信号
	select {
	case <-p.ctx.Done():
		//context被正常取消
		return nil
	default:
		return fmt.Errorf("failed to cancel the context")
	}
}

func (p *PythonInstance) getId() (string, error) {
	if p.Id == "" {
		return "", fmt.Errorf("id is not exists")
	}
	return p.Id, nil
}

type ShellInstance struct {
	*gen.TaskBase
	Id         string
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (s *ShellInstance) execute() error {
	//	通过外部传入的ctx来控制当前shell脚本的执行
	err := kits.RunShellScript(s.ctx, s.TaskBase.ScriptPath, ("shell_job_" + s.Id), "/logs")
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *ShellInstance) stop() error {
	if s.cancelFunc == nil || s.ctx == nil {
		return fmt.Errorf("cancel function or context is not exists")
	}
	//正常生成cancel func，开始发送取消信号
	s.cancelFunc()
	//检查是否获取到context结束信号
	select {
	case <-s.ctx.Done():
		//context被正常取消
		return nil
	default:
		return fmt.Errorf("failed to cancel the context")
	}
}

func (s *ShellInstance) getId() (string, error) {
	if s.Id == "" {
		return "", fmt.Errorf("id is not exists")
	}
	return s.Id, nil
}

// 通过task生成instance

func NewInstance(ctx context.Context, request *gen.ExecuteInstanceRequest) (Instance, error) {
	//获取实例id
	id := kits.GetInstanceID(request.Task.Id, int32(request.Task.Frequency), request.ScheduleTimestamp)
	//继承上级来的context，生成子context用于接受信号
	myCtx, cancel := context.WithCancel(ctx)
	switch request.Task.Type {
	case 0:
		//生成python instance
		result := &PythonInstance{
			TaskBase:          request.Task,
			Id:                id,
			ctx:               myCtx,
			cancelFunc:        cancel,
			pythonInterpreter: "python3",
		}
		return result, nil
	case 1:
		//	生成shell instance
		result := &ShellInstance{
			TaskBase:   request.Task,
			Id:         id,
			ctx:        myCtx,
			cancelFunc: cancel,
		}
		return result, nil
	default:
		cancel()
		return nil, fmt.Errorf("invalid task type")
	}

}
