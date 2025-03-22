package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCheckinRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinRecordByUserIdLogic {
	return &GetCheckinRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCheckinRecordByUserIdLogic) GetCheckinRecordByUserId(in *pb.GetCheckinRecordByUserIdReq) (*pb.GetCheckinRecordByUserIdResp, error) {
	checkinRecord := new(model.CheckinRecord)
	integral := new(model.Integral)

	dbTaskProgress, err := l.svcCtx.TaskProgressModel.FindOneByUserId(l.ctx, in.UserId)
	if errors.Is(err, model.ErrNotFound) {
		dbTaskProgress.UserId = in.UserId
		err := l.svcCtx.TaskProgressModel.Insert(l.ctx, nil, dbTaskProgress)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_INSERT_BY_USER_ID), "Failed to insert task progress : %+v , err: %v", dbTaskProgress, err)
		}
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_FIND_ONE_BY_USER_ID), "Failed to find task progress by user id : %+v , err: %v", in, err)
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		dbRecord, err := l.svcCtx.CheckinRecordModel.FindOneByUserId(l.ctx, in.UserId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_CHECKIN_RECORD_ERROR), "Failed to find checkin record by user id : %+v , err: %v", in, err)
		}
		if errors.Is(err, model.ErrNotFound) {
			dbRecord = &model.CheckinRecord{
				UserId:                in.UserId,
				ContinuousCheckinDays: 0,
				State:                 constants.StateNotCheckin,
			}
			err = l.svcCtx.CheckinRecordModel.Insert(l.ctx, nil, dbRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_INSERT_ERROR), "CheckinRecordModel Insert : %+v , err: %v", dbRecord, err)
			}
		}

		dbIntegral, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_INTEGRAL_ERROR), "IntegralModel FindOneByUserId : %+v , err: %v", dbIntegral, err)
		}
		if errors.Is(err, model.ErrNotFound) {
			dbIntegral = &model.Integral{
				UserId:   in.UserId,
				Integral: 0,
			}
			err = l.svcCtx.IntegralModel.Insert(l.ctx, nil, dbIntegral)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_INSERT_ERROR), "IntegralModel Insert : %+v , err: %v", dbIntegral, err)
			}
		}

		_ = copier.Copy(checkinRecord, dbRecord)
		_ = copier.Copy(integral, dbIntegral)

		today := time.Now().UTC().Truncate(24 * time.Hour)
		targetDay := checkinRecord.LastCheckinDate.Time.UTC().Truncate(24 * time.Hour)

		switch {
		case targetDay.Equal(today):
			return nil
		case targetDay.Equal(today.Add(-24 * time.Hour)):
			if checkinRecord.ContinuousCheckinDays >= 7 {
				checkinRecord.ContinuousCheckinDays = 0
				checkinRecord.State = constants.StateNotCheckin
				err := l.svcCtx.CheckinRecordModel.Update(l.ctx, db, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_UPDATE_ERROR), "Failed to update checkin record : %+v , err: %v", checkinRecord, err)
				}
			} else if checkinRecord.State != constants.StateNotCheckin {
				checkinRecord.State = constants.StateNotCheckin
				err := l.svcCtx.CheckinRecordModel.Update(l.ctx, db, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_UPDATE_ERROR), "Failed to update checkin record : %+v , err: %v", checkinRecord, err)
				}
			}
			return nil
		default:
			// 不需要一直更新数据库，只有有一个不是0时才更新
			if checkinRecord.State != 0 || checkinRecord.ContinuousCheckinDays != 0 {
				checkinRecord.ContinuousCheckinDays = 0
				checkinRecord.State = 0
				err := l.svcCtx.CheckinRecordModel.Update(l.ctx, db, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_CHECKIN_RECORD_UPDATE_ERROR), "Failed to update checkin record : %+v , err: %v", checkinRecord, err)
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_CHECKIN_RECORD_BY_USER_ID_ERROR), "Failed to get checkin record by user id : %+v , err: %v", in, err)
	}
	return &pb.GetCheckinRecordByUserIdResp{
		ContinuousCheckinDays: checkinRecord.ContinuousCheckinDays,
		State:                 checkinRecord.State,
		Integral:              integral.Integral,
		SubStatus:             dbTaskProgress.IsSubCheckin,
	}, nil
}
