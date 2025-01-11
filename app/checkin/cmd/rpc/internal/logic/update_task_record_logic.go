package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskRecordLogic {
	return &UpdateTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskRecordLogic) UpdateTaskRecord(in *pb.UpdateTaskRecordReq) (*pb.UpdateTaskRecordResp, error) {
	dbTask, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_FIND_ONE_ERROR), "TaskModel FindOne error: %v", err)
	}

	dbTaskRecord, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, in.TaskId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_FIND_BY_USER_ID_AND_TASK_ID), "TaskRecordModel FindByUserIdAndTaskId error: %v", err)
	}

	if dbTaskRecord.IsFinished == constants.TaskHasAwarded {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_HAS_AWARDED), "Task has been awarded")
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		dbTaskRecord.IsFinished = constants.TaskHasAwarded
		err := l.svcCtx.TaskRecordModel.Update(l.ctx, db, dbTaskRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_UPDATE_ERROR), "TaskRecordModel Update error: %v", err)
		}

		integralRecord := new(model.IntegralRecord)
		integralRecord.Integral = dbTask.Integral
		integralRecord.Content = "任务奖励"
		integralRecord.UserId = in.UserId
		err = l.svcCtx.IntegralRecordModel.Insert(l.ctx, db, integralRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_RECORD_INSERT_ERROR), "IntegralRecordModel Insert error: %v", err)
		}

		dbIntegral, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_INTEGRAL_ERROR), "IntegralModel FindOneByUserId error: %v", err)
		}

		dbIntegral.Integral += dbTask.Integral
		err = l.svcCtx.IntegralModel.Update(l.ctx, db, dbIntegral)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_INTEGRAL_UPDATE_ERROR), "IntegralModel Update error: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TRANSACTION_ERROR), "TransactCtx error: %v", err)
	}

	return &pb.UpdateTaskRecordResp{}, nil
}
