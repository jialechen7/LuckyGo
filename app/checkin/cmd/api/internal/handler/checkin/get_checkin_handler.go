package checkin

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/logic/checkin"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获得签到状态以及积分
func GetCheckinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCheckinReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		err := validator.New().StructCtx(r.Context(), req)
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		l := checkin.NewGetCheckinLogic(r.Context(), svcCtx)
		resp, err := l.GetCheckin(&req)
		response.HttpResult(r, w, resp, err)
	}
}
