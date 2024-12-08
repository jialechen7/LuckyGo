package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)
var userAuthOmitColumns = []string{"create_time", "update_time"}

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		customUserAuthLogicModel
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}

	customUserAuthLogicModel interface {
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn *gorm.DB, c cache.CacheConf) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn, c),
	}
}

func (m *defaultUserAuthModel) customCacheKeys(data *UserAuth) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
