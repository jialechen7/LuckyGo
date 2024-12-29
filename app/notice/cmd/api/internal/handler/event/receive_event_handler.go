package event

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/logic/event"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 接收小程序回调消息
func ReceiveEventHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReceiveEventReq
		logx.Info("ReceiveEventHandler received an event", r)
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := event.NewReceiveEventLogic(r.Context(), svcCtx)
		resp, err := l.ReceiveEvent(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
