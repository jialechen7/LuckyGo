package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryStatisticLogic {
	return &GetLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryStatisticLogic) GetLotteryStatistic(in *pb.GetLotteryStatisticReq) (*pb.GetLotteryStatisticResp, error) {
	createdCount, err := l.svcCtx.LotteryModel.GetCreatedCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_CREATED_COUNT_BY_USER_ID), "GetCreatedCountByUserId err : %v , userId : %d", err, in.UserId)
	}
	wonCount, err := l.svcCtx.LotteryParticipationModel.GetWonCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_WON_COUNT_BY_USER_ID), "GetWonCountByUserId err : %v , userId : %d", err, in.UserId)
	}
	participationCount, err := l.svcCtx.LotteryParticipationModel.GetParticipationCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATION_COUNT_BY_USER_ID), "GetParticipationCountByUserId err : %v , userId : %d", err, in.UserId)
	}

	return &pb.GetLotteryStatisticResp{
		CreatedCount:       createdCount,
		WonCount:           wonCount,
		ParticipationCount: participationCount,
	}, nil
}
