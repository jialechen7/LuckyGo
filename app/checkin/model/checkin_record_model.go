package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CheckinRecordModel = (*customCheckinRecordModel)(nil)
var checkinRecordOmitColumns = []string{"create_time", "update_time"}

type (
	// CheckinRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCheckinRecordModel.
	CheckinRecordModel interface {
		checkinRecordModel
		customCheckinRecordLogicModel
		FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error)
	}

	customCheckinRecordModel struct {
		*defaultCheckinRecordModel
	}

	customCheckinRecordLogicModel interface {
	}
)

func (c *customCheckinRecordModel) FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error) {
	record := &CheckinRecord{}
	err := c.QueryNoCacheCtx(ctx, &record, func(db *gorm.DB, v interface{}) error {
		return db.Table(c.table).Where("user_id = ?", userId).First(v).Error
	})
	if err != nil {
		return nil, err
	}
	return record, err
}

// NewCheckinRecordModel returns a model for the database table.
func NewCheckinRecordModel(conn *gorm.DB, c cache.CacheConf) CheckinRecordModel {
	return &customCheckinRecordModel{
		defaultCheckinRecordModel: newCheckinRecordModel(conn, c),
	}
}

func (m *defaultCheckinRecordModel) customCacheKeys(data *CheckinRecord) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
