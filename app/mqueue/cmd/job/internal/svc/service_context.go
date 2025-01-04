package svc

import (
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	LotteryRpc  lottery.LotteryZrpcClient
}

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Errorf("【asynq server】 exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		LotteryRpc:  lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
	}
}
