package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLastIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLastIdLogic {
	return &GetCommentLastIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLastIdLogic) GetCommentLastId(in *pb.GetCommentLastIdReq) (*pb.GetCommentLastIdResp, error) {
	id, err := l.svcCtx.CommentModel.GetCommentLastId(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_COMMENT_LAST_ID_ERROR), "model error: %v", err)
	}

	return &pb.GetCommentLastIdResp{
		LastId: id,
	}, nil
}
