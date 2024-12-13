package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUserIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUserIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUserIdsLogic {
	return &GetUserInfoByUserIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByUserIdsLogic) GetUserInfoByUserIds(in *pb.GetUserInfoByUserIdsReq) (*pb.GetUserInfoByUserIdsResp, error) {
	dbUsers, err := l.svcCtx.UserModel.FindUserInfoByUserIds(l.ctx, in.UserIds)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_USERS_INFO_ERROR), "find user info by user ids error: %v", err)
	}

	var list []*pb.UserInfoForComment
	for _, dbUser := range dbUsers {
		user := new(pb.UserInfoForComment)
		_ = copier.Copy(user, dbUser)
		list = append(list, user)
	}

	return &pb.GetUserInfoByUserIdsResp{
		UserInfo: list,
	}, nil
}
