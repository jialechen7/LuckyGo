package user

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)

	pbUserInfo, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserInfo, "GetUserInfo err : %v , userId : %d  , userInfoResp : %+v", err, userId, pbUserInfo)
	}

	userInfo := types.User{}
	_ = copier.Copy(&userInfo, pbUserInfo.User)

	pbStatistic, err := l.svcCtx.LotteryRpc.GetLotteryStatistic(l.ctx, &lottery.GetLotteryStatisticReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetLotteryStatistic, "GetLotteryStatistic err : %v , userId : %d  , userLotteryInfoResp : %+v", err, userId, pbStatistic)
	}
	_ = copier.Copy(&userInfo, pbStatistic)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
