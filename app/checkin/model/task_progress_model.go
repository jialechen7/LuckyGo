package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TaskProgressModel = (*customTaskProgressModel)(nil)
var taskProgressOmitColumns = []string{"create_time", "update_time"}

type (
	// TaskProgressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskProgressModel.
	TaskProgressModel interface {
		taskProgressModel
		customTaskProgressLogicModel

		FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error)
	}

	customTaskProgressModel struct {
		*defaultTaskProgressModel
	}

	customTaskProgressLogicModel interface {
	}
)

func (c *customTaskProgressModel) FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error) {
	taskProgress := &TaskProgress{}
	err := c.QueryNoCacheCtx(ctx, &taskProgress, func(db *gorm.DB, v interface{}) error {
		return db.Table(c.table).Where("user_id = ?", userId).First(v).Error
	})
	if err != nil {
		return nil, err
	}

	return taskProgress, nil
}

// NewTaskProgressModel returns a model for the database table.
func NewTaskProgressModel(conn *gorm.DB, c cache.CacheConf) TaskProgressModel {
	return &customTaskProgressModel{
		defaultTaskProgressModel: newTaskProgressModel(conn, c),
	}
}

func (m *defaultTaskProgressModel) customCacheKeys(data *TaskProgress) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
