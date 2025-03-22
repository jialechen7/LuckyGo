package {{.PkgName}}

import (
	"net/http"

	"github.com/jialechen7/go-lottery/common/response"
	"github.com/go-playground/validator/v10"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r,w,err)
			return
		}

        err := validator.New().StructCtx(r.Context(), req)
        if err != nil {
            response.ParamErrorResult(r, w, err)
            return
        }

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		response.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
