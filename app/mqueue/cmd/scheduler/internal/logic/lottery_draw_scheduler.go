package logic

import (
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/jobtype"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *MqueueScheduler) LotteryDrawScheduler() {
	task := asynq.NewTask(jobtype.ScheduleLotteryDraw, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【LotteryDrawScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	logx.Infof("【LotteryDrawScheduler】 registered an entry: %q \n", entryID)
}
