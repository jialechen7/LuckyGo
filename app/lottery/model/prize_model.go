package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ PrizeModel = (*customPrizeModel)(nil)
var prizeOmitColumns = []string{"create_time", "update_time"}

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		customPrizeLogicModel
		FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error)
		FindFirstLevelPrizeByLotteryId(ctx context.Context, lotteryId int64) (*Prize, error)
		DecrStock(ctx context.Context, prizeId int64) error
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}

	customPrizeLogicModel interface {
	}
)

func (c *customPrizeModel) DecrStock(ctx context.Context, prizeId int64) error {
	return c.ExecCtx(ctx, func(db *gorm.DB) error {
		tx := db.Model(&Prize{}).
			Where("id = ? AND stock > 0", prizeId).
			Update("stock", gorm.Expr("stock - 1"))

		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return ErrRowsAffectedZero
		}
		return nil
	})
}

func (c *customPrizeModel) FindFirstLevelPrizeByLotteryId(ctx context.Context, lotteryId int64) (*Prize, error) {
	var prize Prize
	err := c.QueryNoCacheCtx(ctx, &prize, func(db *gorm.DB, v interface{}) error {
		return db.Where("lottery_id = ? AND level = 1", lotteryId).First(v).Error
	})
	if err != nil {
		return nil, err
	}
	return &prize, nil
}

func (c *customPrizeModel) FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error) {
	var prizes []*Prize
	err := c.QueryNoCacheCtx(ctx, &prizes, func(db *gorm.DB, v interface{}) error {
		return db.Where("lottery_id = ?", lotteryId).Find(v).Error
	})
	if err != nil {
		return nil, err
	}
	return prizes, nil
}

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn *gorm.DB, c cache.CacheConf) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c),
	}
}

func (m *defaultPrizeModel) customCacheKeys(data *Prize) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
