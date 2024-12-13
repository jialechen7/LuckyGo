package lottery

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/logic/lottery"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 抽奖人
func SearchParticipationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchLotteryParticipationReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := lottery.NewSearchParticipationLogic(r.Context(), svcCtx)
		resp, err := l.SearchParticipation(&req)
		response.HttpResult(r, w, resp, err)
	}
}
