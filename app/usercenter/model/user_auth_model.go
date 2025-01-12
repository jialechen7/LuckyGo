package model

import (
	"context"
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

		FindOneByUserId(ctx context.Context, userId int64) (*UserAuth, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}

	customUserAuthLogicModel interface {
	}
)

func (c *customUserAuthModel) FindOneByUserId(ctx context.Context, userId int64) (*UserAuth, error) {
	userAuth := &UserAuth{}
	err := c.QueryNoCacheCtx(ctx, userAuth, func(conn *gorm.DB, v interface{}) error {
		return conn.Where("user_id = ?", userId).First(v).Error
	})
	if err != nil {
		return nil, err
	}
	return userAuth, nil
}

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
