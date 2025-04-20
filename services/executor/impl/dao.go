package impl

import (
	"context"
	"fmt"
	services_conf "github.com/PinkPinkPigg/dora/pkg/config/services"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"gorm.io/gorm"
	"time"
)

//数据库连接
//直接把taskBase当成taskPo吧

// 获取instance的数据
func (s *ServiceImpl) findInstance(ctx context.Context, instanceId string) (*gen.InstanceBase, error) {
	tx := s.db.WithContext(ctx).Table(services_conf.INSTANCE_TABLE_NAME)
	if tx.Error != nil {
		s.l.Errorf("connect to db: %v failed, error:%v", services_conf.INSTANCE_TABLE_NAME, tx.Error)
		return nil, fmt.Errorf("DB connection failed")
	}
	result := &gen.InstanceBase{
		ExecuteTime: 0, //初始化一下重试次数
	}
	tx.Where("id = ?", instanceId).First(result)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("DB found failed")
	}
	return result, nil
}

//改变某个instance的状态

func (s *ServiceImpl) alterInstanceInfo(ctx context.Context, instanceId string, status gen.InstanceStatus) error {
	info, err := s.findInstance(ctx, instanceId)
	if err != nil {
		return err
	}
	if info.Id == "" {
		//	不存在该instance且不是新增状态，报错
		//todo:这里应该统一一个处理错误码的封装，而不是直接返回一个明细错误
		if status != gen.InstanceStatus_WAITING_TO_SUBMIT {
			return fmt.Errorf("instance %s not found", instanceId)
		}
		//	否则自行初始化info
		info.Id = instanceId
		info.CreateTimestamp = uint64(time.Now().Unix())
	}
	//	修改instance的信息为要求信息
	switch status {
	case 1:
		//	修改为执行中，执行次数需要+1
		info.ExecuteTime = info.ExecuteTime + 1
	}
	info.Status = status
	info.ModifyTimestamp = uint64(time.Now().Unix())

	tx := s.db.WithContext(ctx).Table(services_conf.INSTANCE_TABLE_NAME).Save(info)
	if tx.Error != nil {
		s.l.Errorf("alter instance %s to db failed, error:%v", instanceId, tx.Error)
		return fmt.Errorf("DB alter failed")
	}
	return nil
}
