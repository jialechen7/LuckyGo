package comment

import (
	"context"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/comment"
	"github.com/jialechen7/go-lottery/app/comment/model"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"maps"
	"slices"

	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论列表
func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	pbResp, err := l.svcCtx.CommentRpc.SearchComment(l.ctx, &comment.SearchCommentReq{
		LastId: req.LastId,
		Limit:  req.PageSize,
		Sort:   req.Sort,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetCommentList, "rpc error: %v", err)
	}

	userIds := make(map[int64]struct{})
	for _, item := range pbResp.Comment {
		userIds[item.UserId] = struct{}{}
	}

	userInfos, err := l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
		UserIds: slices.Collect(maps.Keys(userIds)),
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserInfos, "rpc error: %v", err)
	}

	userInfoMap := make(map[int64]*types.User)
	for _, userInfo := range userInfos.UserInfo {
		maskedNickName := userInfo.Nickname
		if len(maskedNickName) > 2 {
			maskedNickName = maskedNickName[:1] + "**" + maskedNickName[len(maskedNickName)-1:]
		} else {
			maskedNickName = maskedNickName + "**"
		}
		userInfoMap[userInfo.Id] = &types.User{
			Id:       userInfo.Id,
			Nickname: maskedNickName,
			Avatar:   userInfo.Avatar,
		}
	}

	commentList := make([]types.Comments, 0)
	for _, item := range pbResp.Comment {
		commentItem := new(types.Comments)
		_ = copier.Copy(commentItem, &item)
		commentItem.User = *userInfoMap[item.UserId]
		commentList = append(commentList, *commentItem)
	}

	return &types.CommentListResp{
		List: commentList,
	}, nil
}
