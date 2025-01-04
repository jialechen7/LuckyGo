package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

// 策略模式
type LotteryStrategy interface {
	Run() error
}

type PeopleLotteryStrategy struct {
	l           *AnnounceLotteryLogic
	CurrentTime time.Time
}

type TimeLotteryStrategy struct {
	l           *AnnounceLotteryLogic
	CurrentTime time.Time
}

func (p *TimeLotteryStrategy) Run() error {
	l := p.l
	pendingList, err := l.svcCtx.LotteryModel.GetPendingLotteryListWhichTypeIsTimeStrategyAndAnnouncedTimeBeforeNow(l.ctx, time.Now())
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PENDING_LOTTERY_LIST_FAIL), "GetPendingLotteryListWhichTypeIsTimeStrategyAndAnnouncedTimeBeforeNow err:%v", err)
	}
	for _, lottery := range pendingList {
		// 开奖
		err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			// 获取所有的奖品
			prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindByLotteryId err:%v", err)
			}
			// 获取所有参与者
			participatorIds, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipatorsByLotteryId(l.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATORS_ERROR), "GetLotteryParticipatorsByLotteryId err:%v", err)
			}

			// 开奖逻辑
			winners, err := l.drawLottery(l.ctx, lottery.Id, prizes, participatorIds)
			if err != nil {
				return err
			}

			// 更新lotteryId状态为已开奖
			err = l.svcCtx.LotteryModel.UpdateStatusById(l.ctx, lottery.Id, constants.LotteryHasAnnounced)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_STATUS_ERROR), "UpdateStatusById err:%v", err)
			}

			// 更新参与表，添加中奖者的prizeId
			for _, winner := range winners {
				err := l.svcCtx.LotteryParticipationModel.UpdateWinner(l.ctx, winner.LotteryId, winner.UserId, winner.PrizeId)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_WINNER_ERROR), "UpdateWinner err:%v", err)
				}
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ANNOUNCE_LOTTERY_FAIL), "AnnounceLottery TransactCtx err:%v", err)
		}
	}

	return nil
}

func (p *PeopleLotteryStrategy) Run() error {
	// 获取所有未开奖，且开奖类型是满足人数的
	l := p.l
	pendingList, err := l.svcCtx.LotteryModel.GetPendingLotteryListWhichTypeIsPeopleStrategy(l.ctx)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PENDING_LOTTERY_LIST_FAIL), "GetPendingLotteryListWhichTypeIsPeopleStrategy err:%v", err)
	}
	for _, lottery := range pendingList {
		// 判断是否满足开奖条件：1. 满足人数 2. 开奖时间已经到了（大于等于当前时间）
		participatorsCount, err := l.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(l.ctx, lottery.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATORS_COUNT_FAIL), "GetParticipatorsCountByLotteryId err:%v", err)
		}
		// 开奖时间未到，且参与人数未满足，才不会开奖
		if time.Now().Before(lottery.AnnounceTime) && participatorsCount < lottery.JoinNumber {
			continue
		}
		// 开奖
		err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
			// 获取所有的奖品
			prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PRIZE_ERROR), "FindByLotteryId err:%v", err)
			}
			// 获取所有参与者
			participatorIds, err := l.svcCtx.LotteryParticipationModel.GetLotteryParticipatorsByLotteryId(l.ctx, lottery.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_PARTICIPATORS_ERROR), "GetLotteryParticipatorsByLotteryId err:%v", err)
			}

			// 开奖逻辑
			winners, err := l.drawLottery(l.ctx, lottery.Id, prizes, participatorIds)
			if err != nil {
				return err
			}

			// 更新lotteryId状态为已开奖
			err = l.svcCtx.LotteryModel.UpdateStatusById(l.ctx, lottery.Id, constants.LotteryHasAnnounced)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_STATUS_ERROR), "UpdateStatusById err:%v", err)
			}

			// 更新参与表，添加中奖者的prizeId
			for _, winner := range winners {
				err := l.svcCtx.LotteryParticipationModel.UpdateWinner(l.ctx, winner.LotteryId, winner.UserId, winner.PrizeId)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_WINNER_ERROR), "UpdateWinner err:%v", err)
				}
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ANNOUNCE_LOTTERY_FAIL), "AnnounceLottery TransactCtx err:%v", err)
		}
	}

	return nil
}

/*
Function: drawLottery
Description: 开奖逻辑
Algorithm Goal:
1. 公平性：每个参与者有相同的中奖概率
2. 中奖者数量：中奖者数量不超过奖品数量
3. 奖品稀有度：奖品的稀有度不同，中奖概率不同
4. 高性能：开奖逻辑要求高性能
*/
func (l *AnnounceLotteryLogic) drawLottery(ctx context.Context, lotteryId int64, prizes []*model.Prize, participatorIds []int64) ([]Winner, error) {
	winners := make([]Winner, 0)

	// Fisher–Yates shuffle算法，实现纯随机性，参考链接：https://gaohaoyang.github.io/2016/10/16/shuffle-algorithm/
	for i := len(participatorIds) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		participatorIds[i], participatorIds[j] = participatorIds[j], participatorIds[i]
	}

	// 简易版本算法
	for _, prize := range prizes {
		// 从参与者中随机选取中奖者
		for i := 0; i < int(prize.Count); i++ {
			if len(participatorIds) == 0 {
				break
			}
			idx := rand.Intn(len(participatorIds))
			winner := Winner{
				LotteryId: lotteryId,
				UserId:    participatorIds[idx],
				PrizeId:   prize.Id,
			}
			winners = append(winners, winner)
		}
	}

	return winners, nil
}

func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	var strategy LotteryStrategy
	switch in.AnnounceType {
	case constants.AnnounceTypeTimeLottery:
		strategy = &TimeLotteryStrategy{
			l:           l,
			CurrentTime: time.Now(),
		}
	case constants.AnnounceTypePeopleLottery:
		strategy = &PeopleLotteryStrategy{
			l:           l,
			CurrentTime: time.Now(),
		}
	}
	if strategy == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ANNOUNCE_LOTTERY_FAIL), "AnnounceLottery, invalid announceType:%v", in.AnnounceType)
	}
	err := strategy.Run()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ANNOUNCE_LOTTERY_FAIL), "AnnounceLottery err:%v", err)
	}

	return &pb.AnnounceLotteryResp{}, nil
}
