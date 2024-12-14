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

type GetUserCreatedLotteryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户发起的抽奖列表
func NewGetUserCreatedLotteryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCreatedLotteryListLogic {
	return &GetUserCreatedLotteryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCreatedLotteryListLogic) GetUserCreatedLotteryList(req *types.GetUserCreatedLotteryListReq) (resp *types.GetUserCreatedLotteryListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetUserCreatedList(l.ctx, &pb.GetUserCreatedListReq{
		UserId: utility.GetUserIdFromCtx(l.ctx),
		LastId: req.LastId,
		Size:   req.Size,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserCreatedLottery, "rpc GetUserCreatedLotteryList error: %v", err)
	}

	list := make([]*types.Lottery, 0)
	for _, pbItem := range pbResp.List {
		item := &types.Lottery{}
		_ = copier.Copy(item, pbItem)
		list = append(list, item)
	}

	return &types.GetUserCreatedLotteryListResp{
		List: list,
	}, nil
}
