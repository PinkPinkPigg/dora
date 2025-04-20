package impl

import (
	"context"
	"fmt"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

var client gen.ExecutorServiceClient

//var ctx = context.Background()

func TestServiceImpl_GetInstanceInfo(t *testing.T) {

}

func TestServiceImpl_ExecuteInstance(t *testing.T) {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client = gen.NewExecutorServiceClient(conn)
	myctx := context.Background()
	_, err = client.ExecuteInstance(myctx, &gen.ExecuteInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 900,
		Force:             false,
	})

	_, err = client.ExecuteInstance(myctx, &gen.ExecuteInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 1000,
		Force:             false,
	})

	_, err = client.ExecuteInstance(myctx, &gen.ExecuteInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 8000,
		Force:             false,
	})

	_, err = client.ExecuteInstance(myctx, &gen.ExecuteInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 9000,
		Force:             false,
	})
	//for {
	//	select {
	//	case <-myctx.Done():
	//		fmt.Println("client done")
	//		return
	//
	//	}
	//
	//}
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = client.CancelInstance(myctx, &gen.CancelInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 900,
	})

	time.Sleep(200 * time.Second)

	_, err = client.CancelInstance(myctx, &gen.CancelInstanceRequest{
		Task:              PythonTask2,
		ScheduleTimestamp: 1000,
	})

}
