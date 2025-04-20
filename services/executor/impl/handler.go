package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ServiceImpl struct {
	//继承UnimplementedExecutorServiceServer，同时实现方法
	gen.UnimplementedExecutorServiceServer
	WorkerPool *WorkerPool //	执行任务线程池
	db         *gorm.DB
	l          logger.Logger
}

var ExecutorService *ServiceImpl

func init() {
	//设置日志

	zapConfig := zap.DefaultConfig()
	zapConfig.Level = zap.DebugLevel // 设置日志级别
	zapConfig.ToFiles = true
	// 程序每启动一次，不必都生成一个新的日志文件
	zapConfig.Files.RotateOnStartup = false
	zapConfig.Files.Name = "ExecutorService.log"
	zapConfig.Files.Path = "/logs/"

	err := zap.Configure(zapConfig)

	zap.L().Named("INIT").Info("Logger init success! \n")
	//https://github.com/PinkPinkPigg/DouSheng/blob/main/dou_kit/conf/load.go 代码学习

	ExecutorService = &ServiceImpl{
		UnimplementedExecutorServiceServer: gen.UnimplementedExecutorServiceServer{},
		WorkerPool:                         NewWorkerPool(500), //暂时写死线程池大小
		l:                                  zap.L().Named("executor service"),
	}

	mysql := kits.NewMysql()
	db, err := mysql.GetDB()
	if err != nil {
		ExecutorService.l.Errorf("get db error, %s", err.Error())
	}
	ExecutorService.db = db
	ExecutorService.l.Infof("init executor service success!")
	ExecutorService.WorkerPool.Start() //启动线程池，开始接收instance
}

func GetExecutorService() *ServiceImpl {
	return ExecutorService
}

func (s *ServiceImpl) defaultRecallFunc(ctx context.Context, instanceId string, status gen.InstanceStatus, err error) {
	//	默认回调函数，当异步提交的任务成功或者失败时执行
	err = s.alterInstanceInfo(ctx, instanceId, status)
	if err != nil {
		s.l.Errorf("instance %v recall func execute failed", instanceId)
	}

}

func (s *ServiceImpl) testRecallFunc(ctx context.Context, instanceId string, status gen.InstanceStatus, err error) {
	//	测试回调函数
	fmt.Printf("instance %s test recall func execute, status is %d \n", instanceId, int(status))
}

func (s *ServiceImpl) testSleepRecallFunc(ctx context.Context, instanceId string, status gen.InstanceStatus, err error) {
	//	测试回调函数
	fmt.Printf("instance %s test recall func execute, status is %d \n we will sleep 10seconds", instanceId, int(status))
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(strconv.Itoa(i) + " seconds pass!")
	}
}

func (s *ServiceImpl) ExecuteInstance(ctx context.Context, request *gen.ExecuteInstanceRequest) (*gen.ExecuteInstanceResponse, error) {
	//注意，由于这是个异步任务，ctx传入后会随着接口调用完成和终止，因此不能让该ctx传入一异步instance
	instance, err := NewInstance(request, s.defaultRecallFunc)
	if err != nil {
		return nil, fmt.Errorf("new instance failed: %s", err.Error())
	}
	//实例创建成功，开始加入队列并让协程执行
	err = s.WorkerPool.submitInstance(instance)
	//err = s.alterInstanceInfo(instance.ctx, instanceId, gen.InstanceStatus_WAITING_TO_SUBMIT) //这里不应该由ServiceImpl调用,换成submitInstance时调用recall
	return &gen.ExecuteInstanceResponse{
		Status: 0, //未来处理更多状态码
		Extra:  make(map[string]string),
	}, err
}

func (s *ServiceImpl) CancelInstance(ctx context.Context, request *gen.CancelInstanceRequest) (*gen.CancelInstanceResponse, error) {
	id := kits.GetInstanceID(request.Task.Id, int32(request.Task.Frequency), request.ScheduleTimestamp)
	err := s.WorkerPool.CancelInstance(ctx, id)
	if err != nil {
		s.l.Errorf("cancel instance %s failed, %s", id, err.Error())
		return &gen.CancelInstanceResponse{
				Status: 1,
				Extra:  nil,
			},
			fmt.Errorf("cancel instance failed")
	}
	//	成功终止任务，修改mysql状态
	s.l.Infof("cancel instance %s success", id)
	err = s.alterInstanceInfo(ctx, id, gen.InstanceStatus_CANCEL)
	if err != nil {
		s.l.Infof("alter instance %s status to cancel failed", id)
		return &gen.CancelInstanceResponse{
				Status: 1,
				Extra:  nil,
			},
			fmt.Errorf("cancel instance success, alter instace status to cancel failed")
	}
	return &gen.CancelInstanceResponse{
			Status: 0,
			Extra:  nil,
		},
		nil
}
func (s *ServiceImpl) GetInstanceInfo(ctx context.Context, request *gen.GetInstanceInfoRequest) (*gen.GetInstanceInfoResponse, error) {
	taskId := kits.GetInstanceID(uint64(request.TaskId), int32(request.Frequency), request.ScheduleTimestamp)
	info, err := s.findInstance(ctx, taskId)
	if err != nil {
		return nil, err
	}
	return &gen.GetInstanceInfoResponse{
		Instance: info,
		Extra:    make(map[string]string),
	}, nil
}
