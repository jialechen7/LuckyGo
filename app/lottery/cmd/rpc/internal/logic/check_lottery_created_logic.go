package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLotteryCreatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLotteryCreatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLotteryCreatedLogic {
	return &CheckLotteryCreatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLotteryCreatedLogic) CheckLotteryCreated(in *pb.CheckLotteryCreatedReq) (*pb.CheckLotteryCreatedResp, error) {
	count, err := l.svcCtx.LotteryModel.GetCreatedCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_CREATED_COUNT), "Failed to get lottery created count : %v", err)
	}

	count = min(count, 1)

	return &pb.CheckLotteryCreatedResp{
		Created: count,
	}, nil
}
