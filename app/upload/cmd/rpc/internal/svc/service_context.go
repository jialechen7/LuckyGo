package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/jialechen7/go-lottery/app/upload/cmd/rpc/internal/config"
	"github.com/jialechen7/go-lottery/app/upload/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ServiceContext struct {
	Config          config.Config
	UploadFileModel model.UploadFileModel
	MinioClient     *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	ossClient, _ := minio.New(c.OssConf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.OssConf.AccessKey, c.OssConf.SecretKey, ""),
		Secure: c.OssConf.UseSSL,
	})
	db, _ := mysql.Connect(c.MySQL)
	return &ServiceContext{
		Config:          c,
		UploadFileModel: model.NewUploadFileModel(db, c.Cache),
		MinioClient:     ossClient,
	}
}
