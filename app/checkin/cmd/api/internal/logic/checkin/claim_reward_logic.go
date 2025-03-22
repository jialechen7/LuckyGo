package checkin

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/checkin"
	"github.com/jialechen7/go-lottery/app/checkin/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClaimRewardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 领取任务奖励
func NewClaimRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClaimRewardLogic {
	return &ClaimRewardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClaimRewardLogic) ClaimReward(req *types.ClaimRewardReq) (resp *types.ClaimRewardResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CheckinRpc.UpdateTaskRecord(l.ctx, &checkin.UpdateTaskRecordReq{
		UserId: userId,
		TaskId: req.TaskId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrClaimReward, "UpdateTaskRecordReq error: %v", err)
	}
	return
}
