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

type LotteryListAfterLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录后获取抽奖列表
func NewLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryListAfterLoginLogic {
	return &LotteryListAfterLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryListAfterLoginLogic) LotteryListAfterLogin(req *types.LotteryListReq) (resp *types.LotteryListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetLotteryListAfterLogin(l.ctx, &pb.GetLotteryListAfterLoginReq{
		UserId:     utility.GetUserIdFromCtx(l.ctx),
		LastId:     req.LastId,
		Limit:      req.PageSize,
		IsSelected: req.IsSelected,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrSearchList, "rpc error: %v", err)
	}

	var lotteryList []types.Lottery
	for _, lottery := range pbResp.List {
		var item types.Lottery
		_ = copier.Copy(&item, lottery)
		lotteryList = append(lotteryList, item)
	}

	return &types.LotteryListResp{
		List: lotteryList,
	}, nil
}
