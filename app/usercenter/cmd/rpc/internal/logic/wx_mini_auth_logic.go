package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxMiniAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(in *pb.WxMiniAuthReq) (*pb.WxMiniAuthResp, error) {
	var userId int64
	err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		user := new(model.User)
		user.Nickname = in.Nickname
		user.Avatar = in.Avatar
		err := l.svcCtx.UserModel.Insert(l.ctx, db, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniAuth user insert err:%v,user:%+v", err, user)
		}
		userId = user.Id

		userAuth := new(model.UserAuth)
		userAuth.UserId = userId
		userAuth.AuthType = in.AuthType
		userAuth.AuthKey = in.AuthKey
		err = l.svcCtx.UserAuthModel.Insert(l.ctx, db, userAuth)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniAuth user_auth insert err:%v,userAuth:%+v", err, userAuth)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "WxMiniAuth transact err:%v", err)
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGenerateToken, "GenerateToken userId : %d", userId)
	}

	return &pb.WxMiniAuthResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
