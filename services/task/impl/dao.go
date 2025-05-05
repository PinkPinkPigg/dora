package impl

import (
	"context"
	"errors"
	"fmt"
	services_conf "github.com/PinkPinkPigg/dora/pkg/config/services"
	"github.com/PinkPinkPigg/dora/services/executor/gen"
	"gorm.io/gorm"
	"time"
)

func (s *ServiceImpl) findTask(ctx context.Context, task_id uint64) (*gen.TaskBase, error) {
	tx := s.db.WithContext(ctx).Table(services_conf.TASK_TABLE_NAME)
	if tx.Error != nil {
		s.l.Errorf("connect tor db %v table: %v failed, error:%v", services_conf.TASK_TABLE_NAME, services_conf.TASK_TABLE_NAME, tx.Error)
		return nil, tx.Error
	}
	result := &gen.TaskBase{}
	tx.Where("id = ?", task_id).Find(result)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		s.l.Errorf("find task id %v failed, error:%v", task_id, tx.Error)
		return nil, fmt.Errorf("DB found failed")
	}
	return result, nil
}

func (s *ServiceImpl) createTask(ctx context.Context, task *gen.TaskBase) error {
	if task.Id != 0 {
		return fmt.Errorf("task id %v already exists", task.Id)
	}
	//	task id为0，新建任务
	tx := s.db.WithContext(ctx).Table(services_conf.TASK_TABLE_NAME)
	if tx.Error != nil {
		s.l.Errorf("connect to db %v table: %v failed, error:%v", services_conf.TASK_TABLE_NAME, services_conf.TASK_TABLE_NAME, tx.Error)
		return tx.Error
	}
	task.CreateTimestamp = uint64(time.Now().Unix())
	task.ModifyTimestamp = uint64(time.Now().Unix())
	tx = tx.Create(task)
	if tx.Error != nil {
		s.l.Errorf("create new task in table %v failed, error:%v", services_conf.TASK_TABLE_NAME, tx.Error)
		return tx.Error
	}
	return nil
}

func (s *ServiceImpl) alterTask(ctx context.Context, id uint64, updateContent interface{}) (*gen.TaskBase, error) {
	tx := s.db.WithContext(ctx).Table(services_conf.TASK_TABLE_NAME)
	if tx.Error != nil {
		s.l.Errorf("[alterTask]connect tor db %v table: %v failed, error:%v", services_conf.TASK_TABLE_NAME, services_conf.TASK_TABLE_NAME, tx.Error)
		return nil, tx.Error
	}
	tmp, err := s.findTask(ctx, id)
	if err != nil {
		return nil, err
	}
	if tmp == nil || tmp.Id == 0 {
		//	证明该任务不存在，需要新建
		s.l.Errorf("[alterTask]task id %v not found", id)
		return nil, fmt.Errorf("task id %v not found", id)
	}
	//id存在
	//更新修改时间
	result := &gen.TaskBase{}
	tx = tx.Model(result).Where("id = ?", id).Updates(updateContent).Update("modify_timestamp", uint64(time.Now().Unix()))
	if tx.Error != nil {
		s.l.Errorf("alter task to db failed, error:%v", tx.Error)
		return nil, tx.Error
	}
	return result, nil
}
