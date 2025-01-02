package comment

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/logic/comment"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取评论列表
func GetCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewGetCommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetCommentList(&req)
		response.HttpResult(r, w, resp, err)
	}
}
