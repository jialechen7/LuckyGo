package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCheckinRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCheckinRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCheckinRecordLogic {
	return &UpdateCheckinRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCheckinRecordLogic) UpdateCheckinRecord(in *pb.UpdateCheckinRecordReq) (*pb.UpdateCheckinRecordResp, error) {
	dbRecord, err := l.svcCtx.CheckinRecordModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_CHECKIN_RECORD_ERROR), "CheckinRecordModel FindOneByUserId : %+v , err: %v", dbRecord, err)
	}

	if errors.Is(err, model.ErrNotFound) {
		dbRecord = &model.CheckinRecord{
			UserId:                in.UserId,
			ContinuousCheckinDays: 0,
			State:                 constants.StateNotCheckin,
		}
		err = l.svcCtx.CheckinRecordModel.Insert(l.ctx, nil, dbRecord)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_INSERT_ERROR), "CheckinRecordModel Insert : %+v , err: %v", dbRecord, err)
		}
	}

	dbIntegral, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_INTEGRAL_ERROR), "IntegralModel FindOneByUserId : %+v , err: %v", dbIntegral, err)
	}
	if errors.Is(err, model.ErrNotFound) {
		dbIntegral = &model.Integral{
			UserId:   in.UserId,
			Integral: 0,
		}
		err = l.svcCtx.IntegralModel.Insert(l.ctx, nil, dbIntegral)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_INSERT_ERROR), "IntegralModel Insert : %+v , err: %v", dbIntegral, err)
		}
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		if dbRecord.State == constants.StateHasCheckin && dbRecord.LastCheckinDate.Time.Day() == time.Now().Day() {
			return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_REPEAT_ERROR), "CheckinRecord State : %+v , err: %v", dbRecord, err)
		}

		// 更新积分值
		integrals := calculateCheckinIntegral(dbRecord.ContinuousCheckinDays)
		dbIntegral.Integral += integrals
		err := l.svcCtx.IntegralModel.Update(l.ctx, db, dbIntegral)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_UPDATE_ERROR), "IntegralModel Update : %+v , err: %v", dbIntegral, err)
		}

		dbRecord.ContinuousCheckinDays += 1
		dbRecord.LastCheckinDate.Time = time.Now()
		dbRecord.LastCheckinDate.Valid = true
		dbRecord.State = constants.StateHasCheckin
		err = l.svcCtx.CheckinRecordModel.Update(l.ctx, db, dbRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_UPDATE_ERROR), "CheckinRecordModel Update : %+v , err: %v", dbRecord, err)
		}

		integralRecord := &model.IntegralRecord{}
		integralRecord.Integral = integrals
		integralRecord.Content = "签到"
		integralRecord.UserId = in.UserId
		err = l.svcCtx.IntegralRecordModel.Insert(l.ctx, db, integralRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_RECORD_INSERT_ERROR), "IntegralRecordModel Insert : %+v , err: %v", integralRecord, err)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_TRANSACT_ERROR), "TransactCtx : %+v , err: %v", dbRecord, err)
	}

	return &pb.UpdateCheckinRecordResp{
		ContinuousCheckinDays: dbRecord.ContinuousCheckinDays,
		State:                 dbRecord.State,
		Integral:              dbIntegral.Integral,
	}, nil
}

func calculateCheckinIntegral(continuousCheckinDays int64) int64 {
	var integral int64
	switch continuousCheckinDays {
	case 0, 1:
		integral = 5
	case 2:
		integral = 10
	case 3:
		integral = 15
	case 4:
		integral = 20
	case 5:
		integral = 30
	case 6:
		integral = 40
	default:
		integral = 0
	}
	return integral
}
