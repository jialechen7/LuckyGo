package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrUpload = xerr.NewErrMsg("上传失败")
var ErrBucketNotFound = xerr.NewErrMsg("bucket不存在")
var ErrUploadOss = xerr.NewErrMsg("上传文件到oss失败")
