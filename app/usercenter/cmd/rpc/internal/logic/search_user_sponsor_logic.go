package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserSponsorLogic {
	return &SearchUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserSponsorLogic) SearchUserSponsor(in *pb.SearchUserSponsorReq) (*pb.SearchUserSponsorResp, error) {
	dbUserSponsors, err := l.svcCtx.UserSponsorModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_USER_SPONSOR_LIST_ERROR), "Failed to search user sponsor err : %v , in :%+v", err, in)
	}

	userSponsors := make([]*pb.UserSponsor, 0)
	for _, dbUserSponsor := range dbUserSponsors {
		userSponsor := new(pb.UserSponsor)
		_ = copier.Copy(userSponsor, dbUserSponsor)
		userSponsors = append(userSponsors, userSponsor)
	}

	return &pb.SearchUserSponsorResp{
		UserSponsor: userSponsors,
	}, nil
}
