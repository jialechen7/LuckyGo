package user_sponsor

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/logic/user_sponsor"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改赞助商
func UpdateSponsorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSponsorReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := user_sponsor.NewUpdateSponsorLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSponsor(&req)
		response.HttpResult(r, w, resp, err)
	}
}
