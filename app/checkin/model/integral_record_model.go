package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ IntegralRecordModel = (*customIntegralRecordModel)(nil)
var integralRecordOmitColumns = []string{"create_time", "update_time"}

type (
	// IntegralRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralRecordModel.
	IntegralRecordModel interface {
		integralRecordModel
		customIntegralRecordLogicModel
	}

	customIntegralRecordModel struct {
		*defaultIntegralRecordModel
	}

	customIntegralRecordLogicModel interface {
	}
)

// NewIntegralRecordModel returns a model for the database table.
func NewIntegralRecordModel(conn *gorm.DB, c cache.CacheConf) IntegralRecordModel {
	return &customIntegralRecordModel{
		defaultIntegralRecordModel: newIntegralRecordModel(conn, c),
	}
}

func (m *defaultIntegralRecordModel) customCacheKeys(data *IntegralRecord) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
