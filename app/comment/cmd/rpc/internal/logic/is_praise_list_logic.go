package logic

import (
	"context"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseListLogic {
	return &IsPraiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseListLogic) IsPraiseList(in *pb.IsPraiseListReq) (*pb.IsPraiseListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.IsPraiseListResp{}, nil
}
