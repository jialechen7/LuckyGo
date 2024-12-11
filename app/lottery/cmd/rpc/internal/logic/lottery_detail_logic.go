package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find lottery error: %v", err)
	}

	resp := new(pb.LotteryDetailResp)
	resp.Lottery = new(pb.Lottery)
	_ = copier.Copy(resp.Lottery, lottery)
	resp.Lottery.AnnounceTime = lottery.AnnounceTime.Unix()
	resp.Lottery.PublishTime = lottery.PublishTime.Time.Unix()
	resp.Lottery.AwardDeadline = lottery.AwardDeadline.Unix()
	resp.Lottery.CreateTime = lottery.CreateTime.Unix()
	resp.Lottery.UpdateTime = lottery.UpdateTime.Unix()

	prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(model.ErrFindPrizes, "find prize error: %v", err)
	}

	for _, prize := range prizes {
		item := new(pb.Prize)
		_ = copier.Copy(item, prize)
		resp.Prizes = append(resp.Prizes, item)
	}

	// TODO: Set resp.isParticipated

	return resp, nil
}
