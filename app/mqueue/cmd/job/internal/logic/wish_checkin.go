package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/checkin"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/rpc/notice"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
)

var WishCheckinHandlerFail = xerr.NewErrMsg("WishCheckinHandler ProcessTask fail")

type WishCheckinHandler struct {
	svcCtx *svc.ServiceContext
}

func NewWishCheckinHandler(svcCtx *svc.ServiceContext) *WishCheckinHandler {
	return &WishCheckinHandler{
		svcCtx: svcCtx,
	}
}
func (l *WishCheckinHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// 调用心愿rpc获取数据(用户、奖励、累计天数)
	pbCheckinResp, err := l.svcCtx.CheckinRpc.NoticeWishCheckin(ctx, &checkin.NoticeWishCheckinReq{})
	if err != nil {
		return errors.Wrapf(WishCheckinHandlerFail, "CheckInRpc fail:%v", err)
	}

	// 循环发送通知
	for _, item := range pbCheckinResp.WishCheckinDatas {
		_, err := l.svcCtx.NoticeRpc.NoticeWishSign(ctx, &notice.NoticeWishSignInReq{
			UserId:     item.UserId,
			Reward:     item.Reward,
			Accumulate: item.Accumulate,
		})
		if err != nil {
			return errors.Wrapf(WishCheckinHandlerFail, "user:%d send message fail:%v", item.UserId, err)
		}
	}

	return nil
}
