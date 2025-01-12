package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/svc"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// task
	mux.Handle(jobtype.MsgWxMiniProgramNotifyUser, NewWxMiniProgramNotifyUserHandler(l.svcCtx))

	//schedule job
	mux.Handle(jobtype.ScheduleLotteryDraw, NewLotteryDrawHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleWishCheckin, NewWishCheckinHandler(l.svcCtx))
	return mux
}
