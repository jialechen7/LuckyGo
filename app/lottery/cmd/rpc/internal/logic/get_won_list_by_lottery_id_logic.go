package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"sort"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListByLotteryIdLogic {
	return &GetWonListByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListByLotteryIdLogic) GetWonListByLotteryId(in *pb.GetWonListByLotteryIdReq) (*pb.GetWonListByLotteryIdResp, error) {
	dbList, err := l.svcCtx.LotteryParticipationModel.GetLotteryWinListByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_WIN_LOTTERY_LIST_ERROR), "GetLotteryWinListByLotteryId error: %v", err)
	}

	wonUserIds := make([]int64, 0)
	for _, item := range dbList {
		wonUserIds = append(wonUserIds, item.UserId)
	}

	// 每个奖品对应中奖的用户，可能一对多
	prize2UserMap := make(map[int64][]int64)
	for _, item := range dbList {
		if _, ok := prize2UserMap[item.PrizeId]; !ok {
			prize2UserMap[item.PrizeId] = make([]int64, 0)
		}
		prize2UserMap[item.PrizeId] = append(prize2UserMap[item.PrizeId], item.UserId)
	}

	// 无人中奖
	if len(wonUserIds) == 0 {
		return &pb.GetWonListByLotteryIdResp{}, nil
	}

	pbUserInfos, err := l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
		UserIds: wonUserIds,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserInfo, "GetUserInfoByUserIds rpc error: %v", err)
	}

	for idx, item := range pbUserInfos.UserInfo {
		if len(item.Nickname) > 2 {
			pbUserInfos.UserInfo[idx].Nickname = string(item.Nickname[0]) + "**" + string(item.Nickname[len(item.Nickname)-1])
		} else {
			pbUserInfos.UserInfo[idx].Nickname = string(item.Nickname[0]) + "**"
		}
		pbUserInfos.UserInfo[idx] = item
	}

	userId2UserInfo := make(map[int64]*pb.UserInfo)
	for _, v := range pbUserInfos.UserInfo {
		userId2UserInfo[v.Id] = &pb.UserInfo{
			Id:       v.Id,
			Nickname: []byte(v.Nickname),
			Avatar:   []byte(v.Avatar),
		}
	}

	dbPrizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_LIST_ERROR), "FindByLotteryId error: %v", err)
	}

	// 按照奖品等级排序
	sort.Slice(dbPrizes, func(i, j int) bool {
		return dbPrizes[i].Level < dbPrizes[j].Level
	})

	list := make([]*pb.WonList, 0)
	for _, dbPrize := range dbPrizes {
		prize := new(pb.Prize)
		_ = copier.Copy(prize, dbPrize)

		users := make([]*pb.UserInfo, 0)
		for _, userId := range prize2UserMap[dbPrize.Id] {
			users = append(users, userId2UserInfo[userId])
		}

		list = append(list, &pb.WonList{
			Prize:       prize,
			Users:       users,
			WinnerCount: int64(len(prize2UserMap[dbPrize.Id])),
		})
	}

	return &pb.GetWonListByLotteryIdResp{
		List: list,
	}, nil
}
