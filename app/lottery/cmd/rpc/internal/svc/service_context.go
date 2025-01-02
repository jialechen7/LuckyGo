package svc

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/config"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config                    config.Config
	LotteryModel              model.LotteryModel
	LotteryParticipationModel model.LotteryParticipationModel
	PrizeModel                model.PrizeModel
	UsercenterRpc             usercenter.Usercenter
	TransactCtx               func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                    c,
		LotteryModel:              model.NewLotteryModel(db, c.Cache),
		LotteryParticipationModel: model.NewLotteryParticipationModel(db, c.Cache),
		PrizeModel:                model.NewPrizeModel(db, c.Cache),
		UsercenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		TransactCtx: func(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}
}
