package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserSponsorLogic {
	return &UpdateUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserSponsorLogic) UpdateUserSponsor(in *pb.UpdateUserSponsorReq) (*pb.UpdateUserSponsorResp, error) {
	userSponsor, err := l.svcCtx.UserSponsorModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR_NOT_FOUND), "find user sponsor failed: %v", err)
	}

	if in.Type != 0 {
		userSponsor.Type = in.Type
	}
	if in.AppletType != 0 {
		userSponsor.AppletType = in.AppletType
	}
	if in.IsShow != 0 {
		userSponsor.IsShow = in.IsShow
	}
	if in.Name != "" {
		userSponsor.Name = in.Name
	}
	if in.Desc != "" {
		userSponsor.Desc = in.Desc
	}
	if in.Avatar != "" {
		userSponsor.Avatar = in.Avatar
	}
	if in.QrCode != "" {
		userSponsor.QrCode = in.QrCode
	}
	if in.InputA != "" {
		userSponsor.InputA = in.InputA
	}
	if in.InputB != "" {
		userSponsor.InputB = in.InputB
	}
	err = l.svcCtx.UserSponsorModel.Update(l.ctx, nil, userSponsor)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_USER_SPONSOR_ERROR), "update user sponsor failed: %v", err)
	}

	return &pb.UpdateUserSponsorResp{}, nil
}
