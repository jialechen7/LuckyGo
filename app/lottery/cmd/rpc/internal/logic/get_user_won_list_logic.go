package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWonListLogic {
	return &GetUserWonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserWonListLogic) GetUserWonList(in *pb.GetUserWonListReq) (*pb.GetUserWonListResp, error) {
	dbList, err := l.svcCtx.LotteryParticipationModel.GetUserLotteryWinList(l.ctx, in.UserId, in.LastId, in.Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_USER_WIN_LOTTERY_LIST_ERROR), "GetUserLotteryWinList error: %v", err)
	}

	var list []*pb.UserWonList
	for _, dbItem := range dbList {
		item := new(pb.UserWonList)
		_ = copier.Copy(item, dbItem)
		item.IsWon = true
		item.CreateTime = dbItem.CreateTime.Unix()
		dbPrize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, dbItem.PrizeId)
		if err != nil && !errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindOne error: %v", err)
		}
		item.Prize = new(pb.Prize)
		if !errors.Is(err, model.ErrNotFound) {
			_ = copier.Copy(item.Prize, dbPrize)
		}
		list = append(list, item)
	}

	return &pb.GetUserWonListResp{
		List: list,
	}, nil
}
