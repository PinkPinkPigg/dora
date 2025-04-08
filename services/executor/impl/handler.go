package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
)

type ServiceImpl struct {
	//继承UnimplementedExecutorServiceServer，同时实现方法
	gen.UnimplementedExecutorServiceServer
	WorkerPool *WorkerPool //	执行任务线程池
}

var ExecutorService *ServiceImpl

func init() {
	ExecutorService = &ServiceImpl{
		UnimplementedExecutorServiceServer: gen.UnimplementedExecutorServiceServer{},
		WorkerPool:                         NewWorkerPool(500), //暂时写死线程池大小
	}
	fmt.Println("new ServiceImpl created")
	ExecutorService.WorkerPool.Start() //启动线程池，开始接收instance
}

func GetExecutorService() *ServiceImpl {
	return ExecutorService
}

func (s *ServiceImpl) ExecuteInstance(ctx context.Context, request *gen.ExecuteInstanceRequest) (*gen.ExecuteInstanceResponse, error) {
	instance, err := NewInstance(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("new instance failed: %s", err.Error())
	}
	//实例创建成功，开始加入队列并让协程执行
	err = s.WorkerPool.submitInstance(instance)
	return &gen.ExecuteInstanceResponse{
		Status: 0, //todo未来处理更多状态码
		Extra:  make(map[string]string),
	}, err
}

func (s *ServiceImpl) CancelInstance(ctx context.Context, request *gen.CancelInstanceRequest) (*gen.CancelInstanceResponse, error) {
	id := kits.GetInstanceID(request.Task.Id, int32(request.Task.Frequency), request.ScheduleTimestamp)
	err := s.WorkerPool.CancelInstance(ctx, id)
	if err != nil {
		return &gen.CancelInstanceResponse{
				Status: 1,
				Extra:  nil,
			},
			fmt.Errorf("cancel instance failed: %s", err.Error())
	}
	//	todo：返回终止任务respond
	return &gen.CancelInstanceResponse{
			Status: 0,
			Extra:  nil,
		},
		nil
}
func (s *ServiceImpl) GetInstanceInfo(ctx context.Context, request *gen.GetInstanceInfoRequest) (*gen.GetInstanceInfoResponse, error) {
	_ = kits.GetInstanceID(uint64(request.TaskId), int32(request.Frequency), request.ScheduleTimestamp)
	//	todo：去mysql中读取这个instance id的数据
	panic("implement me")
}
