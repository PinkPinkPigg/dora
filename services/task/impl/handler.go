package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/pkg/config/common"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	scheduler_gen "github.com/PinkPinkPigg/dora/services/scheduler/gen"
	scheduler_impl "github.com/PinkPinkPigg/dora/services/scheduler/impl"
	"github.com/PinkPinkPigg/dora/services/task/gen"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	*gen.UnimplementedTaskServiceServer
	schedulerClient     scheduler_gen.SchedulerServiceClient
	schedulerClientConn *grpc.ClientConn
	db                  *gorm.DB
	l                   logger.Logger
}

var TaskService *ServiceImpl

func init() {
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = zap.DebugLevel // 设置日志级别
	zapConfig.ToFiles = true
	// 程序每启动一次，不必都生成一个新的日志文件
	zapConfig.Files.RotateOnStartup = false
	zapConfig.Files.Name = "ExecutorService.log"
	zapConfig.Files.Path = "/logs/"

	zap.L().Named("INIT").Info("Logger init success! \n")

	TaskService = &ServiceImpl{
		UnimplementedTaskServiceServer: &gen.UnimplementedTaskServiceServer{},
		l:                              zap.L().Named("ServiceImpl"),
	}

	mysql := kits.NewMysql()
	mysql.Host = common.MYSQL_DOCKER_CONTAINER_NAME
	mysql.Port = common.MYSQL_DOCKER_CONTAINER_PORT
	db, err := mysql.GetDB()

	if err != nil {
		TaskService.l.Errorf("get db error, %s", err.Error())
		return
	}

	TaskService.db = db
	//获得scheduler的client
	client, conn, err := scheduler_impl.GetSchedulerServiceClient("prd")

	if err != nil {
		TaskService.l.Errorf("get schduler client error, %s", err.Error())
		defer conn.Close()
		return
	}

	TaskService.schedulerClient = client
	TaskService.schedulerClientConn = conn
	TaskService.l.Info("init task service success!!")

}

func GetServiceImp() *ServiceImpl {
	return TaskService
}

func (s *ServiceImpl) NewTask(ctx context.Context, req *gen.NewTaskRequest) (*gen.NewTaskRespond, error) {
	if req.Task.Id != 0 {
		//保证是新任务
		return nil, status.Error(codes.FailedPrecondition, "task id should be 0")
	}
	err := s.createTask(ctx, req.Task)
	if err != nil {
		return &gen.NewTaskRespond{
			Status: 1,
			TaskId: 0,
			Extra:  nil,
		}, fmt.Errorf("new task error, %s", err)
	}
	return &gen.NewTaskRespond{
		Status: 1,
		TaskId: req.Task.Id,
		Extra:  nil,
	}, nil
}
func (s *ServiceImpl) GetTask(ctx context.Context, req *gen.GetTaskRequest) (*gen.GetTaskRespond, error) {
	task, err := s.findTask(ctx, req.TaskId)
	resp := &gen.GetTaskRespond{}
	if err != nil {
		s.l.Errorf("get task info error, %s", err)
		return nil, status.Error(codes.FailedPrecondition, "task not found")
	}
	resp.Task = task
	return resp, nil
}
func (s *ServiceImpl) AlterTaskStatus(context.Context, *gen.AlterTaskStatusRequest) (*gen.AlterTaskStatusRespond, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterTaskStatus not implemented")
	//	逻辑判断，如果从离线到在线，将任务加入调度器
	//如果是在线到离线，将任务从调度器中删除
	//通过新旧status相减判断情况
}
func (s *ServiceImpl) AlterTaskFrequency(ctx context.Context, req *gen.AlterTaskFrequencyRequest) (*gen.AlterTaskFrequencyRespond, error) {
	task, err := s.findTask(ctx, req.TaskId)
	if err != nil {
		return &gen.AlterTaskFrequencyRespond{
			Status: 1,
			Extra:  nil,
		}, status.Errorf(codes.FailedPrecondition, "find task error, %s", err)
	}
	resp, err := s.schedulerClient.DeleteTask(ctx, &scheduler_gen.DeleteTaskRequest{TaskId: req.TaskId})
	if err != nil || resp.Status != 0 {
		s.l.Errorf("[AlterTaskFrequency] delete task error, %s", err)
		return &gen.AlterTaskFrequencyRespond{
			Status: 1,
			Extra:  nil,
		}, status.Errorf(codes.FailedPrecondition, "delete task error, %s", err)
	}
	resp2, err := s.schedulerClient.AddTask(ctx, &scheduler_gen.AddTaskRequest{Task: task})
	if err != nil || resp2.Status != 0 {
		s.l.Errorf("[AlterTaskFrequency] add task error, %s", err)
		return &gen.AlterTaskFrequencyRespond{
			Status: 1,
			Extra:  nil,
		}, status.Errorf(codes.FailedPrecondition, "add task error, %s", err)
	}
	return &gen.AlterTaskFrequencyRespond{
		Status: 0,
		Extra:  nil,
	}, nil
	//	让调度器删除当前任务
	// 让调度器重新新增任务
	//	修改mysql
}
func (s *ServiceImpl) AlterTaskScriptPath(ctx context.Context, req *gen.AlterTaskScriptPathRequest) (*gen.AlterTaskScriptPathRespond, error) {
	type descriptionAlter struct {
		description string
		id          uint64
	}
	_, err := s.alterTask(ctx, req.TaskId, &descriptionAlter{
		description: req.ScriptPath,
		id:          req.TaskId,
	})
	if err != nil {
		return &gen.AlterTaskScriptPathRespond{
			Status: 1,
		}, err
	}
	return &gen.AlterTaskScriptPathRespond{
		Status: 0,
	}, nil
	//	直接修改mysql
}
func (s *ServiceImpl) AlterTaskDescription(ctx context.Context, req *gen.AlterTaskDescriptionRequest) (*gen.AlterTaskDescriptionRespond, error) {
	type descriptionAlter struct {
		description string
		id          uint64
	}
	_, err := s.alterTask(ctx, req.TaskId, &descriptionAlter{
		description: req.Description,
		id:          req.TaskId,
	})
	if err != nil {
		return &gen.AlterTaskDescriptionRespond{
			Status: 1,
		}, err
	}
	return &gen.AlterTaskDescriptionRespond{
		Status: 0,
	}, nil
}

func (s *ServiceImpl) Close() {
	//负责整体连接的关闭
	s.schedulerClientConn.Close()
}

func (s *ServiceImpl) DeleteTask(ctx context.Context, req *gen.DeleteTaskRequest) (*gen.DeleteTaskRespond, error) {
	task, err := s.findTask(ctx, req.TaskId)
	resp := &gen.DeleteTaskRespond{
		Status: 0,
		Extra:  nil,
	}

	if err != nil || task.Id == 0 {
		resp.Status = 1
		if err == nil {
			err = fmt.Errorf("task not found")
		}
		s.l.Errorf("delete task error, %s", err)
		return resp, status.Error(codes.FailedPrecondition, "task not found")
	}
	//任务存在，开始执行删除
	//调度器中删除任务
	deleteTaskResp, err := s.schedulerClient.DeleteTask(ctx, &scheduler_gen.DeleteTaskRequest{TaskId: req.TaskId})
	if err != nil || deleteTaskResp.Status != 0 {
		s.l.Errorf("delete task error, %s", err)
		resp.Status = 1
		//todo:统一错误处理
		return resp, status.Error(codes.FailedPrecondition, "")
	}
	return resp, nil
	//	让调度器cover executor service内该任务的删除
}
