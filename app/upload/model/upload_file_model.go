package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UploadFileModel = (*customUploadFileModel)(nil)
var uploadFileOmitColumns = []string{"create_time", "update_time"}

type (
	// UploadFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadFileModel.
	UploadFileModel interface {
		uploadFileModel
		customUploadFileLogicModel
	}

	customUploadFileModel struct {
		*defaultUploadFileModel
	}

	customUploadFileLogicModel interface {
	}
)

// NewUploadFileModel returns a model for the database table.
func NewUploadFileModel(conn *gorm.DB, c cache.CacheConf) UploadFileModel {
	return &customUploadFileModel{
		defaultUploadFileModel: newUploadFileModel(conn, c),
	}
}

func (m *defaultUploadFileModel) customCacheKeys(data *UploadFile) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
