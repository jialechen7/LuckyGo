package user_sponsor

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 赞助商详情
func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDetailLogic) SponsorDetail(req *types.SponosorDetailReq) (resp *types.SponosorDetailResp, err error) {
	pbResp, err := l.svcCtx.UsercenterRpc.SponsorDetail(l.ctx, &pb.SponsorDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserSponsorDetail, "rpc error: %v", err)
	}

	resp = new(types.SponosorDetailResp)
	_ = copier.Copy(resp, pbResp)

	return resp, nil
}
