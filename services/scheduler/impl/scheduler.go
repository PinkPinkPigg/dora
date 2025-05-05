package impl

import (
	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	*gocron.Scheduler // 嵌套指针类型

}

var scheduler *Scheduler

func (s *Scheduler) GetActiveJobListFromDB() *[]Job {

}

func GetScheduler() *Scheduler {
	return scheduler
}
