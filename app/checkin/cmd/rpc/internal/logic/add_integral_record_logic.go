package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegralRecordLogic {
	return &AddIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddIntegralRecordLogic) AddIntegralRecord(in *pb.AddIntegralRecordReq) (*pb.AddIntegralRecordResp, error) {
	dbIntegralRecord := &model.IntegralRecord{}
	_ = copier.Copy(dbIntegralRecord, in)
	err := l.svcCtx.IntegralRecordModel.Insert(l.ctx, nil, dbIntegralRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_RECORD_INSERT_ERROR), "IntegralRecordModel Insert : %+v , err: %v", dbIntegralRecord, err)
	}

	return &pb.AddIntegralRecordResp{}, nil
}
