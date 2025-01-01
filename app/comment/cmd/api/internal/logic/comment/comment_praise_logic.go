package comment

import (
	"context"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/types"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/comment"
	"github.com/jialechen7/go-lottery/app/comment/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPraiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 评论点赞/取消点赞
func NewCommentPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPraiseLogic {
	return &CommentPraiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPraiseLogic) CommentPraise(req *types.CommentPraiseReq) (resp *types.CommentPraiseResp, err error) {
	pbReq := &comment.AddPraiseReq{}
	pbReq.CommentId = req.Id
	pbReq.UserId = utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CommentRpc.AddPraise(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrInsertPraise, "CommentPraise rpc error: %v", err)
	}

	return
}
