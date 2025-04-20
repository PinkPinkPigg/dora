package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"testing"
)

func TestServiceImpl_alterInstanceInfo(t *testing.T) {
	service := GetExecutorService()
	err := service.alterInstanceInfo(context.Background(), "1234", gen.InstanceStatus_WAITING_TO_SUBMIT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pass!")
}

func TestServiceImpl_alterInstanceInfo2(t *testing.T) {
	service := GetExecutorService()
	err := service.alterInstanceInfo(context.Background(), "1234", gen.InstanceStatus_EXECUTING)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pass!")
}
