package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/checkin"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/config"
	"github.com/jialechen7/go-lottery/app/notice/cmd/rpc/notice"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	AsynqServer   *asynq.Server
	LotteryRpc    lottery.LotteryZrpcClient
	CheckinRpc    checkin.Checkin
	NoticeRpc     notice.Notice
	WxMiniProgram *miniProgram.MiniProgram
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

func newMiniProgram(c config.Config) *miniProgram.MiniProgram {
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  c.WxMiniConf.AppId,
		Secret: c.WxMiniConf.AppSecret,
		Token:  c.WxMsgConf.EventToken,
		AESKey: c.WxMsgConf.EncodingAESKey,
	})
	logx.Must(err)
	return app
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
		NoticeRpc:     notice.NewNotice(zrpc.MustNewClient(c.NoticeRpcConf)),
		WxMiniProgram: newMiniProgram(c),
	}
}
