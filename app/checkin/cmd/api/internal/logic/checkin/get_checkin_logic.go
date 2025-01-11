package checkin

import (
	"context"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/rpc/checkin"
	"github.com/jialechen7/go-lottery/common/utility"

	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获得签到状态以及积分
func NewGetCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinLogic {
	return &GetCheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCheckinLogic) GetCheckin(req *types.GetCheckinReq) (resp *types.GetCheckinResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbRecord, err := l.svcCtx.CheckinRpc.GetCheckinRecordByUserId(l.ctx, &checkin.GetCheckinRecordByUserIdReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCheckinResp{
		ContinuousCheckinDays: pbRecord.ContinuousCheckinDays,
		State:                 pbRecord.State,
		Integral:              pbRecord.Integral,
		SubStatus:             pbRecord.SubStatus,
	}, nil
}
