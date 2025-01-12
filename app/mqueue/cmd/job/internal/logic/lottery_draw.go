package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/svc"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
)

var drawTypeList = []int64{constants.AnnounceTypeTimeLottery, constants.AnnounceTypePeopleLottery}

var ErrLotteryDrawFail = xerr.NewErrMsg("开奖失败")

type LotteryDrawHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLotteryDrawHandler(svcCtx *svc.ServiceContext) *LotteryDrawHandler {
	return &LotteryDrawHandler{
		svcCtx: svcCtx,
	}
}

func (l *LotteryDrawHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// 用于简单测试定时任务是否正确
	for _, drawType := range drawTypeList {
		_, err := l.svcCtx.LotteryRpc.AnnounceLottery(ctx, &lottery.AnnounceLotteryReq{AnnounceType: drawType})
		if err != nil {
			return errors.Wrapf(ErrLotteryDrawFail, "LotteryDrawHandler announce lottery err:%v", err)
		}
	}

	return nil
}
