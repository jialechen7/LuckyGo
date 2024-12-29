package event

import (
	"net/http"

	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/logic/event"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 验证小程序回调消息
func VerifyEventHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyEventReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := event.NewVerifyEventLogic(r.Context(), svcCtx)
		resp, err := l.VerifyEvent(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
