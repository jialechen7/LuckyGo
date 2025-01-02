package svc

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/config"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	UserModel        model.UserModel
	UserAuthModel    model.UserAuthModel
	UserSponsorModel model.UserSponsorModel
	TransactCtx      func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		UserModel:        model.NewUserModel(db, c.Cache),
		UserAuthModel:    model.NewUserAuthModel(db, c.Cache),
		UserSponsorModel: model.NewUserSponsorModel(db, c.Cache),
		TransactCtx: func(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}
}
