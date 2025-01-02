package user_sponsor

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改赞助商
func NewUpdateSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSponsorLogic {
	return &UpdateSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSponsorLogic) UpdateSponsor(req *types.UpdateSponsorReq) (resp *types.UpdateSponsorResp, err error) {
	pbReq := new(pb.UpdateUserSponsorReq)
	_ = copier.Copy(pbReq, req)
	pbReq.UserId = utility.GetUserIdFromCtx(l.ctx)
	_, err = l.svcCtx.UsercenterRpc.UpdateUserSponsor(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpdateUserSponsor, "rpc error: %v", err)
	}

	return &types.UpdateSponsorResp{}, nil
}
