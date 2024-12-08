package svc

import (
	"github.com/jialechen7/go-lottery/app/upload/cmd/api/internal/config"
	"github.com/jialechen7/go-lottery/app/upload/cmd/rpc/upload"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UploadRpc upload.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UploadRpc: upload.NewUpload(zrpc.MustNewClient(c.UploadRpcConf)),
	}
}
