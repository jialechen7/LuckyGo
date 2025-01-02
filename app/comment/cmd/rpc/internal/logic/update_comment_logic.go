package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/comment/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	dbComment := &model.Comment{}
	_ = copier.Copy(dbComment, in)
	err := l.svcCtx.CommentModel.Update(l.ctx, nil, dbComment)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_COMMENT_ERROR), "UpdateComment rpc error: %v", err)
	}

	return &pb.UpdateCommentResp{}, nil
}
