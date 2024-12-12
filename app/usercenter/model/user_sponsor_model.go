package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserSponsorModel = (*customUserSponsorModel)(nil)
var userSponsorOmitColumns = []string{"create_time", "update_time"}

type (
	// UserSponsorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSponsorModel.
	UserSponsorModel interface {
		userSponsorModel
		customUserSponsorLogicModel
		FindPageByUserId(ctx context.Context, userId, page, limit int64) ([]*UserSponsor, error)
	}

	customUserSponsorModel struct {
		*defaultUserSponsorModel
	}

	customUserSponsorLogicModel interface {
	}
)

func (c *customUserSponsorModel) FindPageByUserId(ctx context.Context, userId, page, limit int64) ([]*UserSponsor, error) {
	var userSponsors []*UserSponsor
	err := c.QueryNoCacheCtx(ctx, &userSponsors, func(conn *gorm.DB, v interface{}) error {
		return conn.Where("user_id = ?", userId).Offset(int((page - 1) * limit)).Limit(int(limit)).Find(v).Error
	})
	if err != nil {
		return nil, err
	}
	return userSponsors, nil
}

// NewUserSponsorModel returns a model for the database table.
func NewUserSponsorModel(conn *gorm.DB, c cache.CacheConf) UserSponsorModel {
	return &customUserSponsorModel{
		defaultUserSponsorModel: newUserSponsorModel(conn, c),
	}
}

func (m *defaultUserSponsorModel) customCacheKeys(data *UserSponsor) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
