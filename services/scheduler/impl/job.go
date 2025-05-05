package impl

import (
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"github.com/go-co-op/gocron"
)

//对gocorn job进行封装

type Job struct {
	*gocron.Job
	*gen.TaskBase
}

func NewJob(base *gen.TaskBase) *Job {
	
}
