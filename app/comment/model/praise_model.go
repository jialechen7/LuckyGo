package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PraiseModel = (*customPraiseModel)(nil)
var praiseOmitColumns = []string{"create_time", "update_time"}

type (
	// PraiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPraiseModel.
	PraiseModel interface {
		praiseModel
		customPraiseLogicModel
	}

	customPraiseModel struct {
		*defaultPraiseModel
	}

	customPraiseLogicModel interface {
	}
)

// NewPraiseModel returns a model for the database table.
func NewPraiseModel(conn *gorm.DB, c cache.CacheConf) PraiseModel {
	return &customPraiseModel{
		defaultPraiseModel: newPraiseModel(conn, c),
	}
}

func (m *defaultPraiseModel) customCacheKeys(data *Praise) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
