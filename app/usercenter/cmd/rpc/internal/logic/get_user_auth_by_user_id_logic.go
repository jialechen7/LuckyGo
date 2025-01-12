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

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserId) (*pb.GetUserAuthByUserIdResp, error) {
	dbUserAuth, err := l.svcCtx.UserAuthModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_USER_AUTH_FIND_ONE_BY_USER_ID_ERROR), "GetUserAuthByUserId FindOneByUserId err:%+v, userId:%d", err, in.UserId)
	}

	pbResp := &pb.UserAuth{}
	_ = copier.Copy(pbResp, dbUserAuth)

	return &pb.GetUserAuthByUserIdResp{
		UserAuth: pbResp,
	}, nil
}
