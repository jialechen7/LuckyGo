package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCreatedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCreatedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCreatedListLogic {
	return &GetUserCreatedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCreatedListLogic) GetUserCreatedList(in *pb.GetUserCreatedListReq) (*pb.GetUserCreatedListResp, error) {
	pbList, err := l.svcCtx.LotteryModel.GetUserCreatedList(l.ctx, in.UserId, in.LastId, in.Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_USER_CREATED_LOTTERY_LIST_ERROR), "GetUserCreatedList error: %v", err)
	}

	list := make([]*pb.Lottery, 0)
	for _, dbItem := range pbList {
		item := new(pb.Lottery)
		_ = copier.Copy(item, dbItem)
		item.PublishTime = dbItem.PublishTime.Time.Unix()
		item.AnnounceTime = dbItem.AnnounceTime.Unix()
		item.AwardDeadline = dbItem.AwardDeadline.Unix()
		item.CreateTime = dbItem.CreateTime.Unix()
		item.UpdateTime = dbItem.UpdateTime.Unix()
		list = append(list, item)
	}

	return &pb.GetUserCreatedListResp{
		List: list,
	}, nil
}
