package svc

import (
	"context"
	"database/sql"
	"github.com/IBM/sarama"
	"github.com/SpectatorNan/gorm-zero/gormc/config/mysql"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/internal/config"
	"github.com/jialechen7/go-lottery/app/lottery/model"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	goredislib "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	Config                    config.Config
	LotteryModel              model.LotteryModel
	LotteryParticipationModel model.LotteryParticipationModel
	PrizeModel                model.PrizeModel
	UsercenterRpc             usercenter.Usercenter
	RedsyncClient             *redsync.Redsync
	RedisClient               *goredislib.Client
	Producer                  sarama.SyncProducer
	TransactCtx               func(context.Context, func(db *gorm.DB) error, ...*sql.TxOptions) error
}

func getAddr(c config.Config) []string {
	return c.Kafka.Host
}

// 创建 kafka 客户端连接的配置
func getConf(c config.Config) *sarama.Config {
	conf := sarama.NewConfig()
	// 生产消息后是否需要通知生产者
	// 同步模式会直接返回
	// 异步模式会返回到Successes和Errors通道中
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Timeout = 5 * time.Second
	conf.Producer.Retry.Max = 5
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	return conf
}

func newKafkaProducer(c config.Config) sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(getAddr(c), getConf(c))
	if err != nil {
		panic(err)
	}
	return producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		panic(err)
	}
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     c.Cache[0].Host,
		Password: c.Cache[0].Pass,
	})
	pool := goredis.NewPool(client)
	svcCtx := &ServiceContext{
		Config:                    c,
		LotteryModel:              model.NewLotteryModel(db, c.Cache),
		LotteryParticipationModel: model.NewLotteryParticipationModel(db, c.Cache),
		PrizeModel:                model.NewPrizeModel(db, c.Cache),
		UsercenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		RedsyncClient:             redsync.New(pool),
		RedisClient:               client,
		Producer:                  newKafkaProducer(c),
		TransactCtx: func(ctx context.Context, fn func(db *gorm.DB) error, opts ...*sql.TxOptions) error {
			return db.WithContext(ctx).Transaction(fn, opts...)
		},
	}

	return svcCtx
}
