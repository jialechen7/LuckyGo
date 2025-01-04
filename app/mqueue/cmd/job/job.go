package main

import (
	"context"
	"flag"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/config"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/logic"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/job.yaml", "Specify the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// 该操作之前是由rest.MustNewServer(c.RestConf)完成的，主要是初始化了prometheus、trace、metricsUrl等
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	logx.DisableStat()

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
