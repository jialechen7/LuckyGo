package svc

import (
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/config"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	LotteryRpc    lottery.LotteryZrpcClient
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
