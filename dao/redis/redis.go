package redis

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			cf.Host,
			cf.Port,
		),
		Password: cf.Password,
		DB:       cf.DB,
		PoolSize: cf.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	if err := rdb.Close(); err != nil {
		zap.L().Fatal("fail to close redis ", zap.Error(err))
	}
}
