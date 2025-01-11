package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ IntegralModel = (*customIntegralModel)(nil)
var integralOmitColumns = []string{"create_time", "update_time"}

type (
	// IntegralModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralModel.
	IntegralModel interface {
		integralModel
		customIntegralLogicModel
		FindOneByUserId(ctx context.Context, userId int64) (*Integral, error)
	}

	customIntegralModel struct {
		*defaultIntegralModel
	}

	customIntegralLogicModel interface {
	}
)

func (c *customIntegralModel) FindOneByUserId(ctx context.Context, userId int64) (*Integral, error) {
	integral := &Integral{}
	err := c.QueryNoCacheCtx(ctx, &integral, func(db *gorm.DB, v interface{}) error {
		return db.Table(c.table).Where("user_id = ?", userId).First(v).Error
	})
	if err != nil {
		return nil, err
	}
	return integral, nil
}

// NewIntegralModel returns a model for the database table.
func NewIntegralModel(conn *gorm.DB, c cache.CacheConf) IntegralModel {
	return &customIntegralModel{
		defaultIntegralModel: newIntegralModel(conn, c),
	}
}

func (m *defaultIntegralModel) customCacheKeys(data *Integral) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
