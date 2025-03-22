package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTasksLogic {
	return &AddTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTasksLogic) AddTasks(in *pb.AddTasksReq) (*pb.AddTasksResp, error) {
	err := l.svcCtx.TasksModel.Insert(l.ctx, nil, &model.Tasks{
		Type:     in.Type,
		Content:  in.Content,
		Integral: in.Integral,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASKS_INSERT_ERROR), "Failed to insert tasks data : %+v , err: %v", in, err)
	}

	return &pb.AddTasksResp{}, nil
}
