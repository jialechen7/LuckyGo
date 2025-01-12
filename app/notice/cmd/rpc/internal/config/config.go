package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisForAsynq     redis.RedisConf
	UsercenterRpcConf zrpc.RpcClientConf
}
