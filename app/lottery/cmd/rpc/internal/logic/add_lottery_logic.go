package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

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

		for _, pbPrize := range in.Prizes {
			prize := new(model.Prize)
			_ = copier.Copy(prize, pbPrize)
			prize.LotteryId = lotteryId
			err := l.svcCtx.PrizeModel.Insert(l.ctx, db, prize)
			if err != nil {
				return errors.Wrapf(model.ErrCreatePrize, "AddLotteryLogic Insert Prize error: %v", err)
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
