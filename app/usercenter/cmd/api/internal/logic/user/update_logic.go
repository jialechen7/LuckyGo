package user

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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户信息
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	userId := utility.GetUserIdFromCtx(l.ctx)
	pbReq := &pb.UpdateUserBaseInfoReq{}
	_ = copier.Copy(pbReq, req)
	pbReq.Id = userId
	_, err = l.svcCtx.UsercenterRpc.UpdateUserBaseInfo(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(model.ErrUpdateUserInfo, "update user info rpc err: %+v , err: %v", req, err)
	}

	return &types.UserUpdateResp{}, nil
}
