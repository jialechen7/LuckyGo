package upload

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"path/filepath"

	"github.com/jialechen7/go-lottery/common/response"

	"github.com/jialechen7/go-lottery/app/upload/cmd/api/internal/logic/upload"
	"github.com/jialechen7/go-lottery/app/upload/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/upload/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 文件上传
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			response.ParamErrorResult(r, w, err)
			return
		}

		ext := filepath.Ext(header.Filename)
		all, err := io.ReadAll(file)
		if err != nil {
			return
		}

		req.FileName = uuid.New().String()
		req.Size = header.Size
		req.Ext = ext
		req.FileData = all

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		response.HttpResult(r, w, resp, err)
	}
}
