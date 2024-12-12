package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserSponsorLogic {
	return &DelUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserSponsorLogic) DelUserSponsor(in *pb.DelUserSponsorReq) (*pb.DelUserSponsorResp, error) {
	err := l.svcCtx.UserSponsorModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete user sponsor failed: %v", err)
	}

	return &pb.DelUserSponsorResp{}, nil
}
