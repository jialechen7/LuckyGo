package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ TasksModel = (*customTasksModel)(nil)
var tasksOmitColumns = []string{"create_time", "update_time"}

type (
	// TasksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTasksModel.
	TasksModel interface {
		tasksModel
		customTasksLogicModel

		FindAllTasks(ctx context.Context) ([]*Tasks, error)
	}

	customTasksModel struct {
		*defaultTasksModel
	}

	customTasksLogicModel interface {
	}
)

func (c *customTasksModel) FindAllTasks(ctx context.Context) ([]*Tasks, error) {
	tasks := make([]*Tasks, 0)
	err := c.QueryNoCacheCtx(ctx, &tasks, func(db *gorm.DB, v interface{}) error {
		return db.Table(c.table).Find(v).Error
	})
	return tasks, err
}

// NewTasksModel returns a model for the database table.
func NewTasksModel(conn *gorm.DB, c cache.CacheConf) TasksModel {
	return &customTasksModel{
		defaultTasksModel: newTasksModel(conn, c),
	}
}

func (m *defaultTasksModel) customCacheKeys(data *Tasks) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
