package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LotteryModel = (*customLotteryModel)(nil)
var lotteryOmitColumns = []string{"create_time", "update_time"}

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		customLotteryLogicModel
		GetLastId(ctx context.Context) (int64, error)
		LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}

	customLotteryLogicModel interface {
	}
)

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn *gorm.DB, c cache.CacheConf) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c),
	}
}

func (m *defaultLotteryModel) customCacheKeys(data *Lottery) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (c *customLotteryModel) GetLastId(ctx context.Context) (int64, error) {
	lottery := Lottery{}
	err := c.QueryNoCacheCtx(ctx, &lottery, func(db *gorm.DB, v interface{}) error {
		err := db.Order("id desc").Limit(1).Find(&lottery).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return lottery.Id, nil
}

func (c *customLotteryModel) LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error) {
	var list []*Lottery
	err := c.QueryNoCacheCtx(ctx, &list, func(db *gorm.DB, v interface{}) error {
		db = db.Where("id > ?", lastId).Where("is_announced = ?", 0)
		if selected != 0 {
			db = db.Where("is_selected = ?", 1)
		}
		err := db.Order("id desc").Limit(int(limit)).Find(&list).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}
