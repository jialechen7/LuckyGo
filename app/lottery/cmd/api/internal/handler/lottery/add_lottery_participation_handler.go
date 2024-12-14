package lottery

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/logic/lottery"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 参与抽奖
func AddLotteryParticipationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddLotteryParticipationReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := lottery.NewAddLotteryParticipationLogic(r.Context(), svcCtx)
		resp, err := l.AddLotteryParticipation(&req)
		response.HttpResult(r, w, resp, err)
	}
}
