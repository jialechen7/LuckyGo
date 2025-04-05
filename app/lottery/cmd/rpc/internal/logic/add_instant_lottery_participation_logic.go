package logic

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
	"strconv"
	"time"
)

type MessageValue struct {
	LotteryId int64  `json:"lottery_id"`
	UserId    int64  `json:"user_id"`
	Timestamp int64  `json:"timestamp"`
	RequestId string `json:"request_id"`
}

type AddInstantLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInstantLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInstantLotteryParticipationLogic {
	return &AddInstantLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddInstantLotteryParticipationLogic) AddInstantLotteryParticipation(in *pb.AddInstantLotteryParticipationReq) (*pb.AddInstantLotteryParticipationResp, error) {
	// 1. 检查参与的抽奖是否为即抽即中类型
	dbLottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_LOTTERY_BY_ID_ERR), "failed to find lottery by id %d", in.LotteryId)
	}
	if dbLottery.AnnounceType != constants.AnnounceTypeInstant {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ANNOUNCE_LOTTERY_FAIL), "lottery %d is not instant lottery", in.LotteryId)
	}

	//// 分布式锁Key = InstantLottery:{UserId}:{LotteryId}
	//mutexKey := constants.InstantLotteryRedisKey + ":" + strconv.Itoa(int(in.UserId)) + ":" + strconv.Itoa(int(in.LotteryId))
	//mutex := l.svcCtx.RedsyncClient.NewMutex(mutexKey)
	//if err := mutex.Lock(); err != nil {
	//	return nil, errors.Wrapf(err, "failed to lock mutex %s", mutexKey)
	//}
	//defer mutex.Unlock()

	// 2. 判断用户已经参与
	lotteryParticipation, err := l.svcCtx.LotteryParticipationModel.FindOneByLotteryIdUserId(l.ctx, in.LotteryId, in.UserId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_PARTICIPATE_LOTTERY), "failed to find lottery participation by lotteryId %d and userId %d", in.LotteryId, in.UserId)
	}

	if lotteryParticipation != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_USER_ALREADY_PARTICIPATE_LOTTERY), "user %d already participate lottery %d", in.UserId, in.LotteryId)
	}

	msgValue := &pb.MessageValue{
		LotteryId: in.LotteryId,
		UserId:    in.UserId,
		Timestamp: time.Now().Unix(),
		RequestId: uuid.New().String(), // 可以考虑换成雪花 ID
	}

	messageValue, _ := proto.Marshal(msgValue)
	msg := &sarama.ProducerMessage{
		Topic: l.svcCtx.Config.Kafka.Topic,
		Key:   sarama.StringEncoder(strconv.Itoa(int(in.LotteryId))),
		Value: sarama.StringEncoder(messageValue),
	}

	go func() {
		partition, offset, err := l.svcCtx.Producer.SendMessage(msg)
		if err != nil {
			l.Logger.Errorf("Async Send Kafka message failed: %v", err)
		} else {
			l.Logger.Infof("Kafka async sent: partition %d, offset %d", partition, offset)
		}
	}()

	return &pb.AddInstantLotteryParticipationResp{
		Id: 0,
	}, nil
}
