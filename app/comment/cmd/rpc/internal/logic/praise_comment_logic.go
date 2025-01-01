package logic

import (
	"context"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PraiseCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPraiseCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PraiseCommentLogic {
	return &PraiseCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PraiseCommentLogic) PraiseComment(in *pb.PraiseCommentReq) (*pb.PraiseCommentResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PraiseCommentResp{}, nil
}
