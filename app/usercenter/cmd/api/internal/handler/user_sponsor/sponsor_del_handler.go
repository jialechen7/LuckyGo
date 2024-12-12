package user_sponsor

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/logic/user_sponsor"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除赞助商
func SponsorDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SponsorDelReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := user_sponsor.NewSponsorDelLogic(r.Context(), svcCtx)
		resp, err := l.SponsorDel(&req)
		response.HttpResult(r, w, resp, err)
	}
}
