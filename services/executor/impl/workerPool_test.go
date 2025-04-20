package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/pkg/kits"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"testing"
	"time"
)

var pool = NewWorkerPool(500)
var ctx, cancel = context.WithCancel(context.Background())
var service = GetExecutorService()
var PythonTask1 = &gen.TaskBase{
	Id:              100,
	Type:            0, //python任务
	Frequency:       1, //day
	Priority:        0,
	Status:          0, //在线
	Description:     "this is a test py instance",
	CreateTimestamp: 0,
	ModifyTimestamp: 0,
	ScriptPath:      "/projects/dora/scripts/py3/test_py3.py",
	RetryConfig:     nil,
}

var PythonTask2 = &gen.TaskBase{
	Id:              200,
	Type:            0, //python任务
	Frequency:       1, //day
	Priority:        0,
	Status:          0, //在线
	Description:     "this is a test py instance",
	CreateTimestamp: 0,
	ModifyTimestamp: 0,
	ScriptPath:      "/projects/dora/scripts/py3/sleep.py",
	RetryConfig:     nil,
}

var PythonInstance1 = &PythonInstance{
	TaskBase:          PythonTask1,
	Id:                kits.GetInstanceID(PythonTask1.Id, int32(PythonTask1.Frequency), uint64(time.Now().Unix())),
	ctx:               context.Background(),
	cancelFunc:        cancel,
	pythonInterpreter: "python3",
	recallFunc:        service.testRecallFunc,
}

var PythonInstance2 = &PythonInstance{
	TaskBase:          PythonTask2,
	Id:                kits.GetInstanceID(PythonTask2.Id, int32(PythonTask2.Frequency), uint64(time.Now().Unix())),
	ctx:               context.Background(),
	cancelFunc:        cancel,
	pythonInterpreter: "python3",
	recallFunc:        service.testRecallFunc,
}

func init() {
	pool.Start()
}

func Test_submitInstance(t *testing.T) {

	//	构造一些任务
	err := pool.submitInstance(PythonInstance1)

	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case <-ctx.Done():
			return

		}
	}

}

func Test_CancelInstance(t *testing.T) {
	//PythonInstance1.recallFunc = service.testSleepRecallFunc
	err := pool.submitInstance(PythonInstance2)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)
	//睡眠2秒后终止该instance
	err = pool.CancelInstance(ctx, PythonInstance2.Id)
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case <-ctx.Done():
			return

		}
	}
}
