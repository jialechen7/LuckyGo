// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/SpectatorNan/gorm-zero/gormc/batchx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheGoLotteryUsercenterUserIdPrefix     = "cache:goLotteryUsercenter:user:id:"
	cacheGoLotteryUsercenterUserMobilePrefix = "cache:goLotteryUsercenter:user:mobile:"
)

type (
	userModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *User) error
		BatchInsert(ctx context.Context, tx *gorm.DB, news []User) error
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		Update(ctx context.Context, tx *gorm.DB, data *User) error
		BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []User) error
		BatchDelete(ctx context.Context, tx *gorm.DB, datas []User) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		// deprecated. recommend add a transaction in service context instead of using this
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultUserModel struct {
		gormc.CachedConn
		table string
	}

	User struct {
		Id               int64     `gorm:"column:id;primary_key"`
		CreateTime       time.Time `gorm:"column:create_time"`
		UpdateTime       time.Time `gorm:"column:update_time"`
		Mobile           string    `gorm:"column:mobile"`            // æ‰‹æœºå·
		Password         string    `gorm:"column:password"`          // å¯†ç 
		Nickname         string    `gorm:"column:nickname"`          // æ˜µç§°
		Sex              int64     `gorm:"column:sex"`               // æ€§åˆ« 0:ç”· 1:å¥³
		Avatar           string    `gorm:"column:avatar"`            // å¤´åƒ
		Info             string    `gorm:"column:info"`              // ç®€ä»‹
		IsAdmin          int64     `gorm:"column:is_admin"`          // æ˜¯å¦ç®¡ç†å‘˜ 0:å¦ 1:æ˜¯
		Signature        string    `gorm:"column:signature"`         // ä¸ªæ€§ç­¾å
		LocationName     string    `gorm:"column:location_name"`     // åœ°å€åç§°
		Longitude        float64   `gorm:"column:longitude"`         // ç»åº¦
		Latitude         float64   `gorm:"column:latitude"`          // çº¬åº¦
		TotalPrize       int64     `gorm:"column:total_prize"`       // ç´¯è®¡å¥–å“
		Fans             int64     `gorm:"column:fans"`              // ç²‰ä¸æ•°é‡
		AllLottery       int64     `gorm:"column:all_lottery"`       // å…¨éƒ¨æŠ½å¥–åŒ…å«æˆ‘å‘èµ·çš„ã€æˆ‘ä¸­å¥–çš„
		InitiationRecord int64     `gorm:"column:initiation_record"` // å‘èµ·æŠ½å¥–è®°å½•
		WinningRecord    int64     `gorm:"column:winning_record"`    // ä¸­å¥–è®°å½•
	}
)

func (User) TableName() string {
	return "`user`"
}

func newUserModel(conn *gorm.DB, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) GetCacheKeys(data *User) []string {
	if data == nil {
		return []string{}
	}
	goLotteryUsercenterUserIdKey := fmt.Sprintf("%s%v", cacheGoLotteryUsercenterUserIdPrefix, data.Id)
	goLotteryUsercenterUserMobileKey := fmt.Sprintf("%s%v", cacheGoLotteryUsercenterUserMobilePrefix, data.Mobile)
	cacheKeys := []string{
		goLotteryUsercenterUserIdKey, goLotteryUsercenterUserMobileKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultUserModel) Insert(ctx context.Context, tx *gorm.DB, data *User) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Omit(userOmitColumns...).Save(&data).Error
	}, m.GetCacheKeys(data)...)
	return err
}
func (m *defaultUserModel) BatchInsert(ctx context.Context, tx *gorm.DB, news []User) error {

	err := batchx.BatchExecCtx(ctx, m, news, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Create(&news).Error
	})

	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	goLotteryUsercenterUserIdKey := fmt.Sprintf("%s%v", cacheGoLotteryUsercenterUserIdPrefix, id)
	var resp User
	err := m.QueryCtx(ctx, &resp, goLotteryUsercenterUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&User{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	goLotteryUsercenterUserMobileKey := fmt.Sprintf("%s%v", cacheGoLotteryUsercenterUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, goLotteryUsercenterUserMobileKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&User{}).Where("`mobile` = ?", mobile).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, tx *gorm.DB, data *User) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && errors.Is(err, ErrNotFound) {
		return err
	}
	clearKeys := append(m.GetCacheKeys(old), m.GetCacheKeys(data)...)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Omit(userOmitColumns...).Save(data).Error
	}, clearKeys...)
	return err
}
func (m *defaultUserModel) BatchUpdate(ctx context.Context, tx *gorm.DB, olds, news []User) error {
	clearData := make([]User, 0, len(olds)+len(news))
	clearData = append(clearData, olds...)
	clearData = append(clearData, news...)
	err := batchx.BatchExecCtx(ctx, m, clearData, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&news).Error
	})

	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&User{}, id).Error
	}, m.GetCacheKeys(data)...)
	return err
}

func (m *defaultUserModel) BatchDelete(ctx context.Context, tx *gorm.DB, datas []User) error {
	err := batchx.BatchExecCtx(ctx, m, datas, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&datas).Error
	})

	return err
}

// deprecated. recommend add a transaction in service context instead of using this
func (m *defaultUserModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoLotteryUsercenterUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&User{}).Where("`id` = ?", primary).Take(v).Error
}
