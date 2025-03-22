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

type GetLotteryListSlowQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListSlowQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListSlowQueryLogic {
	return &GetLotteryListSlowQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListSlowQueryLogic) GetLotteryListSlowQuery(in *pb.GetLotteryListSlowQueryReq) (*pb.GetLotteryListSlowQueryResp, error) {
	list, err := l.svcCtx.LotteryModel.LotteryListSlowQuery(l.ctx, in.Limit, in.Page, in.IsSelected)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_LIST_ERROR), "get lottery list error: %v", err)
	}

	var pbList []*pb.Lottery
	for _, lottery := range list {
		pbLottery := &pb.Lottery{}
		_ = copier.Copy(pbLottery, lottery)
		pbLottery.PublishTime = lottery.PublishTime.Time.Unix()
		pbLottery.AwardDeadline = lottery.AwardDeadline.Unix()
		pbLottery.AnnounceTime = lottery.AnnounceTime.Unix()
		pbList = append(pbList, pbLottery)
	}

	return &pb.GetLotteryListSlowQueryResp{
		List: pbList,
	}, nil
}
