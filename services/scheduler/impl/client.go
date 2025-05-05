package impl

import (
	"errors"
	"github.com/PinkPinkPigg/dora/pkg/config/common"
	"github.com/PinkPinkPigg/dora/services/scheduler/gen"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetSchedulerServiceClient(environment string) (gen.SchedulerServiceClient, *grpc.ClientConn, error) {
	//environment:环境
	var err error
	var conn *grpc.ClientConn
	switch environment {
	case "prd":
		conn, err = grpc.NewClient(common.SERVICES_SCHEDULER_CONTAINER_NAME+":"+common.SERVICES_SCHEDULER_CONTAINER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case "dev":
		conn, err = grpc.NewClient(common.TEST_SERVICES_SCHEDULER_CONTAINER_NAME+":"+common.TEST_SERVICES_SCHEDULER_CONTAINER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	default:
		zap.L().Error("unsupported environment", zap.String("environment", environment))
		return nil, nil, errors.New("invalid environment")
	}
	if err != nil {
		zap.L().Error("get service client fail", zap.String("environment", environment), zap.Error(err))
		defer conn.Close()
		return nil, nil, err
	}
	schedulerServiceClient := gen.NewSchedulerServiceClient(conn)

	return schedulerServiceClient, conn, nil

}
