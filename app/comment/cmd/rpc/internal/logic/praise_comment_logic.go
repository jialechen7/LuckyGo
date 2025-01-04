package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/comment/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

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
	dbPraise, err := l.svcCtx.PraiseModel.IsPraise(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment error: %v", err)
	}
	// 若点赞，取消点赞
	if dbPraise.Id != 0 {
		err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			err := l.svcCtx.PraiseModel.Delete(l.ctx, db, dbPraise.Id)
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.AddPraiseCount(l.ctx, in.CommentId, -1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment error: %v", err)
		}
	} else {
		err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			err := l.svcCtx.PraiseModel.Insert(l.ctx, db, &model.Praise{
				CommentId: in.CommentId,
				UserId:    in.UserId,
			})
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.AddPraiseCount(l.ctx, in.CommentId, 1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PRAISE_COMMENT_ERROR), "PraiseComment error: %v", err)
		}

	}

	return &pb.PraiseCommentResp{}, nil
}
