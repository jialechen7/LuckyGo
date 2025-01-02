package model

import (
	"context"
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
		IsPraise(ctx context.Context, commentId int64, userId int64) (*Praise, error)
		IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error)
	}

	customPraiseModel struct {
		*defaultPraiseModel
	}

	customPraiseLogicModel interface {
	}
)

func (c *customPraiseModel) IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error) {
	likeList := make([]int64, 0)
	err := c.QueryNoCacheCtx(ctx, &likeList, func(db *gorm.DB, v interface{}) error {
		return db.Table(c.table).Where("comment_id IN (?) AND user_id = ?", commentIds, userId).Pluck("comment_id", v).Error
	})
	if err != nil {
		return nil, err
	}

	return likeList, nil
}

func (c *customPraiseModel) IsPraise(ctx context.Context, commentId int64, userId int64) (*Praise, error) {
	dbPraise := &Praise{}
	err := c.QueryNoCacheCtx(ctx, &dbPraise, func(db *gorm.DB, v interface{}) error {
		return db.Where("comment_id = ? AND user_id = ?", commentId, userId).Limit(1).Find(v).Error
	})
	if err != nil {
		return nil, err
	}

	return dbPraise, nil
}

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
