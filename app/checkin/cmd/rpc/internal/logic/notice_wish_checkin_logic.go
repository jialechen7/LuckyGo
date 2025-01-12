package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeWishCheckinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishCheckinLogic {
	return &NoticeWishCheckinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeWishCheckinLogic) NoticeWishCheckin(in *pb.NoticeWishCheckinReq) (*pb.NoticeWishCheckinResp, error) {
	userIds, err := l.svcCtx.TaskProgressModel.FindAllSubscribeUserIds(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_TASK_PROGRESS_FIND_ALL_SUBSCRIBE_USER_IDS_ERROR), "Failed to find all subscribe user ids, err: %v", err)
	}

	list := make([]*pb.NoticeWishCheckinData, 0)
	getCheckinRecordLogic := NewGetCheckinRecordByUserIdLogic(l.ctx, l.svcCtx)
	for _, userId := range userIds {
		pbResp, err := getCheckinRecordLogic.GetCheckinRecordByUserId(&pb.GetCheckinRecordByUserIdReq{
			UserId: userId,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_CHECKIN_RECORD_BY_USER_ID_ERROR), "Failed to get checkin record by user id : %v, err: %v", userId, err)
		}

		if pbResp.State == constants.UserTodayHasCheckin {
			continue
		}
		var accumulate int64
		var award int64
		if pbResp.ContinuousCheckinDays >= 7 {
			accumulate = 1
		} else {
			accumulate = pbResp.ContinuousCheckinDays + 1
		}
		switch accumulate {
		case 1, 2:
			award = 5
		case 3:
			award = 10
		case 4:
			award = 15
		case 5:
			award = 20
		case 6:
			award = 30
		case 7:
			award = 40
		default:
			award = 0
		}
		userRecord := &pb.NoticeWishCheckinData{
			UserId:     userId,
			Accumulate: accumulate,
			Reward:     award,
		}
		list = append(list, userRecord)
	}

	return &pb.NoticeWishCheckinResp{
		WishCheckinDatas: list,
	}, nil
}
