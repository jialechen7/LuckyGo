package event

import (
	"context"
	"net/http"

	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 接收小程序回调消息
func NewReceiveEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveEventLogic {
	return &ReceiveEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveEventLogic) ReceiveEvent(_ *types.ReceiveEventReq, r *http.Request) (resp *types.ReceiveEventResp, err error) {
	// todo: add your logic here and delete this line
	logx.Info("ReceiveEventLogic received an event", r)

	return
}
