package model

import (
	"context"
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
		FindUserInfoByUserIds(ctx context.Context, ids []int64) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}

	customUserLogicModel interface {
	}
)

func (c *customUserModel) FindUserInfoByUserIds(ctx context.Context, ids []int64) ([]*User, error) {
	var users []*User
	err := c.QueryNoCacheCtx(ctx, &users, func(conn *gorm.DB, v interface{}) error {
		return conn.Where("id in (?)", ids).Find(v).Error
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

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
