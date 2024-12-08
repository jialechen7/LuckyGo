package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserModel = (*customUserModel)(nil)
var userOmitColumns = []string{"create_time", "update_time"}

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		customUserLogicModel
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserLogicModel interface {
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn *gorm.DB, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) customCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
