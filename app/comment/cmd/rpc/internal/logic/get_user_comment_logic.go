package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCommentLogic {
	return &GetUserCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCommentLogic) GetUserComment(in *pb.GetUserCommentReq) (*pb.GetUserCommentResp, error) {
	dbList, err := l.svcCtx.CommentModel.GetCommentListByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_COMMENT_LIST_BY_USER_ID_ERROR), "CommentModel GetCommentListByUserId fail , req : %+v , err : %v", in, err)
	}

	commentIds := make([]int64, 0)
	for _, dbComment := range dbList {
		commentIds = append(commentIds, dbComment.Id)
	}

	if len(commentIds) == 0 {
		return &pb.GetUserCommentResp{}, nil
	}

	isPraiseList, err := l.svcCtx.PraiseModel.IsPraiseList(l.ctx, commentIds, in.UserId)
	if err != nil {
		return nil, err
	}

	likeMap := make(map[int64]int64)
	for _, commentId := range commentIds {
		likeMap[commentId] = 0
	}
	for _, commentId := range isPraiseList {
		likeMap[commentId] = 1
	}

	commentList := make([]*pb.Comment, 0)
	for _, dbComment := range dbList {
		pbComment := &pb.Comment{}
		_ = copier.Copy(pbComment, dbComment)
		pbComment.CreateTime = dbComment.CreateTime.Unix()
		pbComment.UpdateTime = dbComment.UpdateTime.Unix()
		pbComment.IsPraise = likeMap[dbComment.Id]

		commentList = append(commentList, pbComment)
	}

	return &pb.GetUserCommentResp{
		Comment: commentList,
	}, nil
}
