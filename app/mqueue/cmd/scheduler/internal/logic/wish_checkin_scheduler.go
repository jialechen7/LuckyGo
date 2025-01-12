package logic

import (
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/jobtype"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *MqueueScheduler) WishCheckinScheduler() {
	task := asynq.NewTask(jobtype.ScheduleWishCheckin, nil)
	// 测试阶段：每分钟执行
	// 实际上线阶段：每天10点执行
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【WishCheckinScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	logx.Infof("【WishCheckinScheduler】 registered an entry: %q \n", entryID)
}
