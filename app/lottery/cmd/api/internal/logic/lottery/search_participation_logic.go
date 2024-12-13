package lottery

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 抽奖人
func NewSearchParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParticipationLogic {
	return &SearchParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParticipationLogic) SearchParticipation(req *types.SearchLotteryParticipationReq) (resp *types.SearchLotteryParticipationResp, err error) {
	pbResp, err := l.svcCtx.LotteryRpc.SearchLotteryParticipation(l.ctx, &pb.SearchLotteryParticipationReq{
		LotteryId: req.LotteryId,
		PageSize:  req.PageSize,
		PageIndex: req.PageIndex,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLotteryParticipation, "SearchLotteryParticipation error: %v", err)
	}

	userIds := make([]int64, 0)
	for _, pbItem := range pbResp.List {
		userIds = append(userIds, pbItem.UserId)
	}

	pbUserInfos, err := l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}

	for idx, item := range pbUserInfos.UserInfo {
		if len(item.Nickname) > 2 {
			item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
		} else {
			item.Nickname = item.Nickname[:] + "**"
		}
		pbUserInfos.UserInfo[idx] = item
	}

	resp = new(types.SearchLotteryParticipationResp)
	_ = copier.Copy(&resp.List, pbUserInfos.UserInfo)
	resp.Count = pbResp.Count

	return resp, nil
}
