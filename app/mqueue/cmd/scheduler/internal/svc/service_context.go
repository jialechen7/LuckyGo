package svc

import (
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/scheduler/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ServiceContext struct {
	Config    config.Config
	Scheduler *asynq.Scheduler
}

func newScheduler(c config.Config) *asynq.Scheduler {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
	}, &asynq.SchedulerOpts{
		Location: location,
		PostEnqueueFunc: func(info *asynq.TaskInfo, err error) {
			if err != nil {
				logx.Errorf("【Scheduler PostEnqueueFunc】 <<<<<<<== =>>>>>> err : %+v , info : %+v", err, info)
			} else {
				logx.Infof("【Scheduler PostEnqueueFunc】 <<<<<<<== =>>>>>> info : %+v", info)
			}
		},
	})
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: newScheduler(c),
	}
}
