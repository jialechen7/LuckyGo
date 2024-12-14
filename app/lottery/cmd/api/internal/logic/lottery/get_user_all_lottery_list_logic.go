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

type GetUserAllLotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户所有抽奖列表
func NewGetUserAllLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllLotteryListLogic {
	return &GetUserAllLotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAllLotteryListLogic) GetUserAllLotteryList(req *types.GetUserAllLotteryListReq) (resp *types.GetUserAllLotteryListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetUserAllList(l.ctx, &pb.GetUserAllListReq{
		UserId: utility.GetUserIdFromCtx(l.ctx),
		LastId: req.LastId,
		Size:   req.Size,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserAllLottery, "rpc GetUserAllLotteryList error: %v", err)
	}
	list := make([]*types.UserLotteryList, 0)
	for _, pbItem := range pbResp.List {
		item := &types.UserLotteryList{}
		if pbItem.IsWon {
			item.IsWon = 1
		}
		item.Prize = new(types.LotteryPrize)
		_ = copier.Copy(item.Prize, pbItem.Prize)
		_ = copier.Copy(item, pbItem)
		list = append(list, item)
	}
	return &types.GetUserAllLotteryListResp{
		List: list,
	}, nil

	return
}
