package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPraiseLogic {
	return &DelPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPraiseLogic) DelPraise(in *pb.DelPraiseReq) (*pb.DelPraiseResp, error) {
	err := l.svcCtx.PraiseModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_DELETE_PRAISE_ERROR), "Failed to delete praise, PraiseModel Delete fail , req : %+v , err : %v", in, err)
	}

	return &pb.DelPraiseResp{}, nil
}
