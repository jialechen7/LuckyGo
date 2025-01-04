package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/scheduler/internal/svc"
)

type MqueueScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqueueScheduler {
	return &MqueueScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqueueScheduler) Register() {
	l.LotteryDrawScheduler()
}
