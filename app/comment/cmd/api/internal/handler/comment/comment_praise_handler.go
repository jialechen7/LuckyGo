package comment

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/logic/comment"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 评论点赞/取消点赞
func CommentPraiseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentPraiseReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := comment.NewCommentPraiseLogic(r.Context(), svcCtx)
		resp, err := l.CommentPraise(&req)
		response.HttpResult(r, w, resp, err)
	}
}
