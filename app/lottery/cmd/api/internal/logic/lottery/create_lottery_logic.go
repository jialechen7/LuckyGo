package lottery

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发起抽奖
func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	logx.Error(req)
	userId := utility.GetUserIdFromCtx(l.ctx)

	var pbPrizes []*pb.Prize
	for _, reqPrize := range req.Prizes {
		pbPrize := new(pb.Prize)
		_ = copier.Copy(pbPrize, reqPrize)
		pbPrizes = append(pbPrizes, pbPrize)
	}

	pbClockTask := new(pb.ClockTask)
	if req.IsClocked == 1 && req.ClockTask != nil {
		_ = copier.Copy(pbClockTask, req.ClockTask)
	} else {
		pbClockTask = nil
	}

	pbReq := new(pb.AddLotteryReq)
	_ = copier.Copy(pbReq, req)
	pbReq.Prizes = pbPrizes
	pbReq.ClockTask = pbClockTask
	pbReq.UserId = userId
	pbResp, err := l.svcCtx.LotteryRpc.AddLottery(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrCreateLottery, "rpc error: %v", err)
	}

	return &types.CreateLotteryResp{
		Id: pbResp.Id,
	}, nil
}
