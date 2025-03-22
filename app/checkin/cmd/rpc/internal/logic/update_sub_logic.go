package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSubLogic) UpdateSub(in *pb.UpdateSubReq) (*pb.UpdateSubResp, error) {
	dbTaskProgress, err := l.svcCtx.TaskProgressModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_FIND_ONE_BY_USER_ID), "TaskProgress Model FindOneByUserId failed: %v", err)
	}

	dbTaskProgress.IsSubCheckin = in.State
	err = l.svcCtx.TaskProgressModel.Update(l.ctx, nil, dbTaskProgress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_UPDATE_BY_USER_ID), "TaskProgress Model Update failed: %v", err)
	}
	if in.State == constants.UserHasSubscribed {
		_, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, constants.UserSubscribeTaskId)
		if errors.Is(err, model.ErrNotFound) {
			addTaskRecordLogic := NewAddTaskRecordLogic(l.ctx, l.svcCtx)
			_, err = addTaskRecordLogic.AddTaskRecord(&pb.AddTaskRecordReq{
				UserId: in.UserId,
				TaskId: constants.UserSubscribeTaskId,
			})
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ADD_TASK_RECORD_ERROR), "AddTaskRecordLogic AddTaskRecord failed: %v", err)
			}
		} else if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_FIND_BY_USER_ID_AND_TASK_ID), "TaskRecord Model FindByUserIdAndTaskId failed: %v", err)
		}

	}

	return &pb.UpdateSubResp{}, nil
}
