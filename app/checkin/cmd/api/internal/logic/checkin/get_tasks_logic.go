package checkin

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/checkin"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获得用户任务完成状态
func NewGetTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksLogic {
	return &GetTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTasksLogic) GetTasks(req *types.GetTasksReq) (resp *types.GetTasksResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	// 查询用户任务进度，返回具体数量
	count, err := l.svcCtx.CheckinRpc.GetTaskProgress(l.ctx, &checkin.GetTaskProgressReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	// 查询用户完成的任务
	tasks, err := l.svcCtx.CheckinRpc.GetTaskRecordByUserId(l.ctx, &checkin.GetTaskRecordByUserIdReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	var taskList []*types.Tasks
	_ = copier.Copy(&taskList, tasks.TaskList)
	for i, task := range taskList {
		switch task.Id {
		case 4:
			task.Count = count.DayCount
			task.NeedCount = 3
		case 7:
			task.Count = count.WeekCount
			task.NeedCount = 30
		default:
			task.Count = -1
			task.NeedCount = -1
		}
		taskList[i] = task
	}
	// 返回任务进度具体数量
	return &types.GetTasksResp{
		TasksList: taskList,
	}, nil
}
