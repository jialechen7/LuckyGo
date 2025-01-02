package comment

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/logic/comment"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除评论
func DeleteCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentDelReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewDeleteCommentLogic(r.Context(), svcCtx)
		resp, err := l.DeleteComment(&req)
		response.HttpResult(r, w, resp, err)
	}
}
