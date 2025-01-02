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
		LotteryList(ctx context.Context, limit, isSelected, lastId int64) ([]*Lottery, error)
		GetLotteryListAfterLogin(ctx context.Context, limit, isSelected, lastId int64, lotteryIds []int64) ([]*Lottery, error)
		GetCreatedCountByUserId(ctx context.Context, userId int64) (int64, error)
		GetUserCreatedList(ctx context.Context, userId, lastId, limit int64) ([]*Lottery, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}

	customLotteryLogicModel interface {
	}
)

func (c *customLotteryModel) GetUserCreatedList(ctx context.Context, userId, lastId, limit int64) ([]*Lottery, error) {
	list := make([]*Lottery, 0)
	err := c.QueryNoCacheCtx(ctx, &list, func(db *gorm.DB, v interface{}) error {
		if lastId > 0 {
			return db.Table(c.table).Where("user_id = ? AND id < ?", userId, lastId).Order("id DESC").Limit(int(limit)).Find(v).Error
		}
		return db.Table(c.table).Where("user_id = ?", userId).Order("id DESC").Limit(int(limit)).Find(v).Error
	})
	if err != nil {
		return []*Lottery{}, err
	}
	return list, err
}

func (c *customLotteryModel) GetCreatedCountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	err := c.QueryNoCacheCtx(ctx, &count, func(db *gorm.DB, v interface{}) error {
		err := db.Model(&Lottery{}).Where("user_id = ?", userId).Count(&count).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn *gorm.DB, c cache.CacheConf) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c),
	}
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

func (c *customLotteryModel) LotteryList(ctx context.Context, limit, isSelected, lastId int64) ([]*Lottery, error) {
	var list []*Lottery
	err := c.QueryNoCacheCtx(ctx, &list, func(db *gorm.DB, v interface{}) error {
		db = db.Where("id < ?", lastId).Where("is_announced = ?", 0)
		if isSelected != 0 {
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

func (c *customLotteryModel) GetLotteryListAfterLogin(ctx context.Context, limit, isSelected, lastId int64, lotteryIds []int64) ([]*Lottery, error) {
	if len(lotteryIds) == 0 {
		list, err := c.LotteryList(ctx, limit, isSelected, lastId)
		if err != nil {
			return nil, err
		}
		return list, nil
	}

	var list []*Lottery
	err := c.QueryNoCacheCtx(ctx, &list, func(db *gorm.DB, v interface{}) error {
		db = db.Where("id < ?", lastId).Where("is_announced = ?", 0).Not(lotteryIds)
		if isSelected != 0 {
			db = db.Where("is_selected = ?", 1)
		}
		err := db.Order("id desc").Limit(int(limit)).Find(v).Error
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

func (m *defaultLotteryModel) customCacheKeys(data *Lottery) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
