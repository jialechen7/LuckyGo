package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryParticipationLogic {
	return &SearchLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryParticipationLogic) SearchLotteryParticipation(in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {
	dbLotteryParticipationList, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipationListByLotteryId(l.ctx, in.LotteryId, in.PageIndex, in.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATED_BY_LOTTERY_ID), "get lottery participation list by lottery id error: %v", err)
	}

	resp := new(pb.SearchLotteryParticipationResp)
	_ = copier.Copy(&resp.List, dbLotteryParticipationList)

	count, err := l.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATIONS_COUNT), "get participators count by lottery id error: %v", err)
	}
	resp.Count = count

	return resp, nil
}
