package logic

import (
	"context"
	"encoding/json"

	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"time"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisPrize struct {
	LotteryId        int64 `json:"lottery_id"`
	Id               int64 `json:"id"`
	Count            int64 `json:"count"`
	Stock            int64 `json:"stock"`
	Level            int64 `json:"level"`
	LeftProbability  int64 `json:"left_probability"`
	RightProbability int64 `json:"right_probability"`
}

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	logx.Error(in)
	var lotteryId int64
	err := l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		lottery := &model.Lottery{
			UserId:        in.UserId,
			Name:          in.Name,
			Thumb:         in.Thumb,
			AwardDeadline: time.Unix(in.AwardDeadline, 0),
			AnnounceTime:  time.Unix(in.AnnounceTime, 0),
			AnnounceType:  in.AnnounceType,
			Introduce:     in.Introduce,
			JoinNumber:    in.JoinNumber,
			SponsorId:     in.SponsorId,
			IsClocked:     in.IsClocked,
		}

		if in.PublishType == constants.PublishTypeNormal {
			lottery.PublishTime.Time = time.Now()
			lottery.PublishTime.Valid = true
		}

		err := l.svcCtx.LotteryModel.Insert(l.ctx, db, lottery)
		if err != nil {
			return errors.Wrapf(model.ErrCreateLottery, "AddLotteryLogic Insert Lottery error: %v", err)
		}
		lotteryId = lottery.Id

		prizePoolKey := constants.InstantLotteryPrizePoolRedisKey + strconv.Itoa(int(lotteryId))
		for _, pbPrize := range in.Prizes {
			prize := new(model.Prize)
			_ = copier.Copy(prize, pbPrize)
			prize.LotteryId = lotteryId
			prize.Stock = prize.Count
      
			err := l.svcCtx.PrizeModel.Insert(l.ctx, db, prize)
			if err != nil {
				return errors.Wrapf(model.ErrCreatePrize, "AddLotteryLogic Insert Prize error: %v", err)
			}

			redisPrize := &RedisPrize{
				LotteryId:        lotteryId,
				Id:               prize.Id,
				Count:            prize.Count,
				Stock:            prize.Stock,
				Level:            prize.Level,
				LeftProbability:  int64(constants.ProbabilityMap[int(prize.Level)][0]),
				RightProbability: int64(constants.ProbabilityMap[int(prize.Level)][1]),
			}
			jsonData, _ := json.Marshal(redisPrize)
			err = l.svcCtx.RedisClient.HSet(l.ctx, prizePoolKey, strconv.FormatInt(redisPrize.Id, 10), jsonData).Err()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.REDIS_ERROR), "AddLotteryLogic Redis HSet error: %v", err)
			}
		}

		// TODO: Add ClockTask

		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddLotteryLogic TransactCtx error: %v", err)
	}

	return &pb.AddLotteryResp{
		Id: lotteryId,
	}, nil
}
