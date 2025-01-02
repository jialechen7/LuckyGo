package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentReq) (*pb.DelCommentResp, error) {
	err := l.svcCtx.CommentModel.Delete(l.ctx, nil, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_DELETE_COMMENT_ERROR), "Failed to delete comment, CommentModel Delete fail , req : %+v , err : %v", in, err)
	}

	return &pb.DelCommentResp{}, nil
}
