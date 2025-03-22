package svc

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/config"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config              config.Config
	CheckinRecordModel  model.CheckinRecordModel
	IntegralModel       model.IntegralModel
	IntegralRecordModel model.IntegralRecordModel
	TasksModel          model.TasksModel
	TaskRecordModel     model.TaskRecordModel
	TaskProgressModel   model.TaskProgressModel
	LotteryRpc          lottery.LotteryZrpcClient
	TransactCtx         func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:              c,
		CheckinRecordModel:  model.NewCheckinRecordModel(db, c.Cache),
		IntegralModel:       model.NewIntegralModel(db, c.Cache),
		IntegralRecordModel: model.NewIntegralRecordModel(db, c.Cache),
		TasksModel:          model.NewTasksModel(db, c.Cache),
		TaskRecordModel:     model.NewTaskRecordModel(db, c.Cache),
		TaskProgressModel:   model.NewTaskProgressModel(db, c.Cache),
		LotteryRpc:          lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		TransactCtx: func(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}
}
