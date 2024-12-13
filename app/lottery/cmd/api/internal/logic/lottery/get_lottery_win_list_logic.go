package lottery

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryWinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前抽奖中奖者名单
func NewGetLotteryWinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinListLogic {
	return &GetLotteryWinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinListLogic) GetLotteryWinList(req *types.GetLotteryWinListReq) (resp *types.GetLotteryWinListResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.GetWonListByLotteryId(l.ctx, &pb.GetWonListByLotteryIdReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLotteryWinList, "rpc GetWonListByLotteryId error: %v", err)
	}

	var list []*types.WonList
	for _, pbItem := range pbResp.List {
		item := &types.WonList{}
		item.Prize = new(types.LotteryPrize)
		_ = copier.Copy(item.Prize, pbItem.Prize)
		item.WinnerCount = pbItem.WinnerCount
		item.Users = make([]*types.UserInfo, 0)
		for _, pbUser := range pbItem.Users {
			user := &types.UserInfo{}
			_ = copier.Copy(user, pbUser)
			item.Users = append(item.Users, user)
		}
		list = append(list, item)
	}

	return &types.GetLotteryWinListResp{
		List: list,
	}, nil
}
