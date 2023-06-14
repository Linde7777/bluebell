package redis

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port,
		),
		Password: settings.Conf.RedisConfig.Password,
		DB:       settings.Conf.RedisConfig.DB,
		PoolSize: settings.Conf.RedisConfig.PoolSize,
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	if err := rdb.Close(); err != nil {
		zap.L().Fatal("fail to close redis ", zap.Error(err))
	}
}
