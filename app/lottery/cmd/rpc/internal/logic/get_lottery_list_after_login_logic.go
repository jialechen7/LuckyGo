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

type GetLotteryListAfterLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListAfterLoginLogic {
	return &GetLotteryListAfterLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListAfterLoginLogic) GetLotteryListAfterLogin(in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	participatedLotteryIds, err := l.svcCtx.LotteryParticipationModel.GetParticipationLotteryIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if in.LastId == 0 {
		id, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GETLASTID_ERROR), "get last id error: %v", err)
		}
		in.LastId = id + 1
	}

	list, err := l.svcCtx.LotteryModel.GetLotteryListAfterLogin(l.ctx, in.Limit, in.IsSelected, in.LastId, participatedLotteryIds)
	if err != nil {
		return nil, err
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

	return &pb.GetLotteryListAfterLoginResp{
		List: pbList,
	}, nil
}
