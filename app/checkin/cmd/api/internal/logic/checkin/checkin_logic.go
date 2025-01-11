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

type CheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 签到操作
func NewCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckinLogic {
	return &CheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckinLogic) Checkin(req *types.CheckinReq) (resp *types.CheckinResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbRecord, err := l.svcCtx.CheckinRpc.UpdateCheckinRecord(l.ctx, &checkin.UpdateCheckinRecordReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrCheckin, "Checkin failed, req: %+v", req)
	}

	return &types.CheckinResp{
		ContinuousCheckinDays: pbRecord.ContinuousCheckinDays,
		State:                 pbRecord.State,
		Integral:              pbRecord.Integral,
	}, nil
}
