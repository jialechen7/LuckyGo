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

type GetUserLotteryWinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户中奖列表
func NewGetUserLotteryWinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLotteryWinListLogic {
	return &GetUserLotteryWinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLotteryWinListLogic) GetUserLotteryWinList(req *types.GetUserLotteryWinListReq) (resp *types.GetUserLotteryWinListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetUserWonList(l.ctx, &pb.GetUserWonListReq{
		UserId: utility.GetUserIdFromCtx(l.ctx),
		LastId: req.LastId,
		Size:   req.Size,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLotteryWinList, "rpc GetUserWonList error: %v", err)
	}
	var list []*types.UserWonList
	for _, pbItem := range pbResp.List {
		item := &types.UserWonList{}
		item.Prize = new(types.LotteryPrize)
		_ = copier.Copy(item.Prize, pbItem.Prize)
		_ = copier.Copy(item, pbItem)
		list = append(list, item)
	}
	return &types.GetUserLotteryWinListResp{
		List: list,
	}, nil
}
