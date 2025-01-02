package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAllListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAllListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAllListLogic {
	return &GetUserAllListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAllListLogic) GetUserAllList(in *pb.GetUserAllListReq) (*pb.GetUserAllListResp, error) {
	dbList, err := l.svcCtx.LotteryParticipationModel.GetAllLotteryListByUserId(l.ctx, in.UserId, in.LastId, in.Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_ALL_USER_LOTTERY_LIST_ERROR), "GetAllLotteryListByUserId error: %v", err)
	}

	list := make([]*pb.UserLotteryList, 0)
	for _, dbItem := range dbList {
		item := new(pb.UserLotteryList)
		_ = copier.Copy(item, dbItem)
		item.CreateTime = dbItem.CreateTime.Unix()
		switch dbItem.IsWon {
		case constants.PrizeNotWon:
			// 若没中奖，返回当前抽奖的一等奖作为封面
			dbPrize, err := l.svcCtx.PrizeModel.FindFirstLevelPrizeByLotteryId(l.ctx, dbItem.LotteryId)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindFirstLevelPrizeByLotteryId error: %v", err)
			}
			item.Prize = new(pb.Prize)
			item.IsWon = false
			_ = copier.Copy(item.Prize, dbPrize)
		case constants.PrizeHasWon:
			dbPrize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, dbItem.PrizeId)
			if err != nil && !errors.Is(err, model.ErrNotFound) {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindOne error: %v", err)
			}
			item.Prize = new(pb.Prize)
			item.IsWon = true
			if !errors.Is(err, model.ErrNotFound) {
				_ = copier.Copy(item.Prize, dbPrize)
			}
		}
		list = append(list, item)
	}

	return &pb.GetUserAllListResp{
		List: list,
	}, nil
}
