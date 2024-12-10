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

type SearchLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryLogic {
	return &SearchLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryLogic) SearchLottery(in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	//if in.LastId == 0 {
	//	id, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
	//	if err != nil {
	//		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GETLASTID_ERROR), "get last id error: %v", err)
	//	}
	//	in.LastId = id + 1
	//}
	list, err := l.svcCtx.LotteryModel.LotteryList(l.ctx, in.Limit, in.IsSelected, in.LastId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_LIST_ERROR), "get lottery list error: %v", err)
	}

	var pbList []*pb.Lottery
	for _, lottery := range list {
		pbLottery := &pb.Lottery{}
		_ = copier.Copy(pbLottery, lottery)
		pbLottery.PublishTime = lottery.PublishTime.Unix()
		pbLottery.AwardDeadline = lottery.AwardDeadline.Unix()
		pbLottery.AnnounceTime = lottery.AnnounceTime.Unix()
		pbList = append(pbList, pbLottery)
	}
	return &pb.SearchLotteryResp{
		Lottery: pbList,
	}, nil
}
