package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBaseInfoLogic {
	return &UpdateUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserBaseInfoLogic) UpdateUserBaseInfo(in *pb.UpdateUserBaseInfoReq) (*pb.UpdateUserBaseInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err: %v", err)
	}
	if in.Nickname != "" {
		user.Nickname = in.Nickname
	}
	if in.Avatar != "" {
		user.Avatar = in.Avatar
	}
	if in.Sex != user.Sex {
		user.Sex = in.Sex
	}
	if in.Info != "" {
		user.Info = in.Info
	}
	if in.Signature != "" {
		user.Signature = in.Signature
	}
	if in.Longitude != 0 {
		user.Longitude = in.Longitude
	}
	if in.Latitude != 0 {
		user.Latitude = in.Latitude
	}
	err = l.svcCtx.UserModel.Insert(l.ctx, nil, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err: %v", err)
	}
	return &pb.UpdateUserBaseInfoResp{}, nil
}
