package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskRecordLogic {
	return &AddTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTaskRecordLogic) AddTaskRecord(in *pb.AddTaskRecordReq) (*pb.AddTaskRecordResp, error) {
	task, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_FIND_ONE_ERROR), "TaskModel FindOne error: %v", err)
	}
	taskRecord := new(model.TaskRecord)
	taskRecord.TaskId = in.TaskId
	taskRecord.UserId = in.UserId
	taskRecord.Type = task.Type
	taskRecord.IsFinished = constants.TaskIsFinished
	err = l.svcCtx.TaskRecordModel.Insert(l.ctx, nil, taskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_INSERT_ERROR), "TaskRecordModel Insert error: %v", err)
	}
	return &pb.AddTaskRecordResp{}, nil
}
