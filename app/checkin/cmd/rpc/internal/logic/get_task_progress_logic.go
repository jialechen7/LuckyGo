package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/lottery"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskProgressLogic {
	return &GetTaskProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type TaskStrategy interface {
	Run() error
}

// NewbieTaskStrategy 新手任务
type NewbieTaskStrategy struct {
	l        *GetTaskProgressLogic
	Logic    *AddTaskRecordLogic
	TaskId   int64
	UserId   int64
	db       *gorm.DB
	Progress *model.TaskProgress
}

func (n *NewbieTaskStrategy) Run() error {
	dbTaskRecord, err := n.l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(n.l.ctx, n.UserId, n.TaskId)
	if errors.Is(err, model.ErrNotFound) {
		switch n.TaskId {
		case constants.NewbieParticipateTaskId:
			pbResp, err := n.l.svcCtx.LotteryRpc.CheckLotteryParticipated(n.l.ctx, &lottery.CheckLotteryParticipatedReq{
				UserId: n.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NEWBIE_CHECK_LOTTERY_PARTICIPATED), "Failed to check lottery participated : %+v , err: %v", pbResp, err)
			}
			if pbResp.Participated == constants.TaskProgressHasParticipatedLottery {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: n.UserId,
					TaskId: n.TaskId,
				}
				_, err := n.Logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ADD_TASK_RECORD_ERROR), "Failed to add task record : %+v , err: %v", addTaskRecord, err)
				}
				// 修改task_progress记录
				n.Progress.IsParticipatedLottery = constants.TaskProgressHasParticipatedLottery
				err = n.l.svcCtx.TaskProgressModel.Update(n.l.ctx, n.db, n.Progress)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_UPDATE_BY_USER_ID), "Failed to update taskProgress data : %+v , err: %v", n.Progress, err)
				}
			}
		case constants.NewbieCreateLotteryTaskId:
			pbResp, err := n.l.svcCtx.LotteryRpc.CheckLotteryCreated(n.l.ctx, &lottery.CheckLotteryCreatedReq{
				UserId: n.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NEWBIE_CHECK_LOTTERY_CREATED), "Failed to check lottery created : %+v , err: %v", pbResp, err)
			}

			if pbResp.Created == constants.TaskProgressHasCreateLottery {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: n.UserId,
					TaskId: n.TaskId,
				}
				_, err := n.Logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ADD_TASK_RECORD_ERROR), "Failed to add task record : %+v , err: %v", addTaskRecord, err)
				}
				// 修改task_progress记录
				n.Progress.IsInitiatedLottery = constants.TaskProgressHasCreateLottery
				err = n.l.svcCtx.TaskProgressModel.Update(n.l.ctx, n.db, n.Progress)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_UPDATE_BY_USER_ID), "Failed to update taskProgress data : %+v , err: %v", n.Progress, err)
				}
			}
		}

	} else if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_RECORD_FIND_BY_USER_ID_AND_TASK_ID), "Failed to find taskRecord data : %+v , err: %v", dbTaskRecord, err)
	}
	return nil
}

func (l *GetTaskProgressLogic) GetTaskProgress(in *pb.GetTaskProgressReq) (*pb.GetTaskProgressResp, error) {
	dbTaskProgress, err := l.svcCtx.TaskProgressModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_FIND_ONE_BY_USER_ID), "Failed to find taskProgress data : %+v , err: %v", dbTaskProgress, err)
	}

	if errors.Is(err, model.ErrNotFound) {
		dbTaskProgress = new(model.TaskProgress)
		dbTaskProgress.UserId = in.UserId
		err := l.svcCtx.TaskProgressModel.Insert(l.ctx, nil, dbTaskProgress)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_INSERT_BY_USER_ID), "Failed to insert taskProgress data : %+v , err: %v", dbTaskProgress, err)
		}
	}

	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		if dbTaskProgress.IsParticipatedLottery != constants.TaskProgressHasParticipatedLottery {
			// 新手任务
			newbieTask := NewbieTaskStrategy{
				l:        l,
				Logic:    NewAddTaskRecordLogic(l.ctx, l.svcCtx),
				TaskId:   constants.NewbieParticipateTaskId,
				UserId:   in.UserId,
				db:       db,
				Progress: dbTaskProgress,
			}
			err := newbieTask.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NEWBIE_RUN), "Failed to run newbie task : %+v , err: %v", newbieTask, err)
			}
		}
		if dbTaskProgress.IsInitiatedLottery != constants.TaskProgressHasCreateLottery {
			// 新手任务
			newbieTask := NewbieTaskStrategy{
				l:        l,
				Logic:    NewAddTaskRecordLogic(l.ctx, l.svcCtx),
				TaskId:   constants.NewbieCreateLotteryTaskId,
				UserId:   in.UserId,
				db:       db,
				Progress: dbTaskProgress,
			}
			err := newbieTask.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NEWBIE_RUN), "Failed to run newbie task : %+v , err: %v", newbieTask, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_TASK_PROGRESS_TRANSACT_ERROR), "Failed to transact taskProgress data : %+v , err: %v", dbTaskProgress, err)
	}

	return &pb.GetTaskProgressResp{}, nil
}
