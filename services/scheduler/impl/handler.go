package impl

import (
	"context"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	executor "github.com/PinkPinkPigg/dora/services/executor/gen"
	"github.com/PinkPinkPigg/dora/services/scheduler/config"
	"github.com/PinkPinkPigg/dora/services/scheduler/gen"
	"github.com/go-co-op/gocron"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
	"time"
)

type ServiceImpl struct {
	*gen.UnimplementedSchedulerServiceServer
	executorClient *executor.ExecutorServiceClient
	scheduler      *Scheduler
	db             *gorm.DB
	l              logger.Logger
}

var impl *ServiceImpl

func init() {
	impl = &ServiceImpl{}
	location, err := time.LoadLocation(config.TIME_AREA)
	impl.l = zap.L().Named("scheduler")
	if err != nil {
		impl.l.Fatal("load location error", zap.Error(err))
		return
	}
	scheduler = &Scheduler{
		Scheduler: gocron.NewScheduler(location),
	}
	impl.scheduler = scheduler
	newMysql := kits.NewMysql()
	db, err := newMysql.GetDB()
	if err != nil {
		impl.l.Fatal("get db error, %s", err.Error())
		return
	}
	impl.db = db
	impl.l.Info("db init success!!")
	//初始化execuotr client

	impl.l.Info("scheduler init success!!")
	//	todo:将所有的活跃任务加入scheduler内
}

func GetServiceImpl() *ServiceImpl {
	return impl
}

func (s *ServiceImpl) AddTask(context.Context, *gen.AddTaskRequest) (*gen.AddTaskRespond, error) {
	//	由任务管理系统吊用，将单个任务加入schedule中
	s.scheduler.Update()

}
func (s *ServiceImpl) DeleteTask(context.Context, *gen.DeleteTaskRequest) (*gen.DeleteTaskRespond, error) {
	panic("implement me")
}
func (s *ServiceImpl) GetAllTaskInScheduler(context.Context, *gen.GetAllTaskInSchedulerRequest) (*gen.GetAllTaskInSchedulerRespond, error) {
	panic("implement me")
}
