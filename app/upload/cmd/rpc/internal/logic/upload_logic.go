package logic

import (
	"bytes"
	"context"
	"github.com/jialechen7/go-lottery/app/upload/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/upload/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/upload/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *pb.FileUploadReq) (*pb.FileUploadResp, error) {
	uploadFile := new(model.UploadFile)
	_ = copier.Copy(uploadFile, in)
	err := l.svcCtx.MinioClient.MakeBucket(l.ctx,
		l.svcCtx.Config.OssConf.Bucket, minio.MakeBucketOptions{Region: l.svcCtx.Config.OssConf.Region})
	if err != nil {
		exists, errBucketExists := l.svcCtx.MinioClient.BucketExists(l.ctx, l.svcCtx.Config.OssConf.Bucket)
		if errBucketExists == nil && exists {
			logx.Infof("We already own %s\n", l.svcCtx.Config.OssConf.Bucket)
		} else {
			logx.Errorf("Failed to create %s\n", l.svcCtx.Config.OssConf.Bucket)
			return nil, errors.Wrapf(model.ErrBucketNotFound, "create bucket with oss client err: %+v , err: %v", in, err)
		}
	} else {
		logx.Infof("Successfully created %s\n", l.svcCtx.Config.OssConf.Bucket)
	}

	reader := bytes.NewReader(in.FileData)
	fileName := in.FileName + in.Ext
	uploadInfo, err := l.svcCtx.MinioClient.PutObject(l.ctx, l.svcCtx.Config.OssConf.Bucket, fileName, reader, in.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return nil, errors.Wrapf(model.ErrUploadOss, "put object with oss client err: %+v , err: %v", in, err)
	}
	url := l.svcCtx.Config.OssConf.Endpoint + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key
	if l.svcCtx.Config.OssConf.UseSSL {
		url = "https://" + url
	} else {
		url = "http://" + url
	}
	uploadFile.Url = url
	err = l.svcCtx.UploadFileModel.Insert(l.ctx, nil, uploadFile)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "save file info with err: %+v , err: %v", in, err)
	}

	return &pb.FileUploadResp{
		Url: url,
	}, nil
}
