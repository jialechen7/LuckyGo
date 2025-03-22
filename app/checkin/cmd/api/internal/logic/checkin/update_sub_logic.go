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

type UpdateSubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 改变订阅签到状态
func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubLogic) UpdateSub(req *types.UpdateSubReq) (resp *types.UpdateSubResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.CheckinRpc.UpdateSub(l.ctx, &checkin.UpdateSubReq{
		UserId: userId,
		State:  req.State,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpdateSub, "UpdateSubReq error: %v", err)
	}
	return
}
