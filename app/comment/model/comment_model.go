package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CommentModel = (*customCommentModel)(nil)
var commentOmitColumns = []string{"create_time", "update_time"}

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		customCommentLogicModel

		GetCommentLastId(ctx context.Context) (int64, error)
		GetCommentList(ctx context.Context, limit, lastId, sort int64) ([]*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}

	customCommentLogicModel interface {
	}
)

func (c *customCommentModel) GetCommentList(ctx context.Context, limit, lastId, sort int64) ([]*Comment, error) {
	commentList := make([]*Comment, 0)
	err := c.QueryNoCacheCtx(ctx, &commentList, func(db *gorm.DB, v interface{}) error {
		db = db.Where("id < ?", lastId)
		if sort == 0 {
			db = db.Order("id desc")
		} else {
			db = db.Order("praise_count desc")
		}
		return db.Limit(int(limit)).Find(v).Error
	})
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func (c *customCommentModel) GetCommentLastId(ctx context.Context) (int64, error) {
	comment := &Comment{}
	err := c.QueryNoCacheCtx(ctx, comment, func(db *gorm.DB, v interface{}) error {
		return db.Order("id desc").Limit(1).Find(v).Error
	})
	if err != nil {
		return 0, err
	}

	return comment.Id, nil
}

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn *gorm.DB, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}

func (m *defaultCommentModel) customCacheKeys(data *Comment) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
