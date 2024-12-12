package user_sponsor

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的赞助商列表
func NewSponsorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorListLogic {
	return &SponsorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorListLogic) SponsorList(req *types.SponsorListReq) (resp *types.SponsorListResp, err error) {
	pbResp, err := l.svcCtx.UsercenterRpc.SearchUserSponsor(l.ctx, &pb.SearchUserSponsorReq{
		Page:   req.Page,
		Limit:  req.PageSize,
		UserId: utility.GetUserIdFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	var sponsorList []types.Sponsor
	for _, pbSponsor := range pbResp.UserSponsor {
		sponsor := types.Sponsor{}
		_ = copier.Copy(&sponsor, pbSponsor)
		sponsorList = append(sponsorList, sponsor)
	}

	return &types.SponsorListResp{
		List: sponsorList,
	}, nil
}
