package logic

import (
	"context"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPraiseLogic {
	return &SearchPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPraiseLogic) SearchPraise(in *pb.SearchPraiseReq) (*pb.SearchPraiseResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchPraiseResp{}, nil
}
