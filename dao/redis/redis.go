package redis

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()
	return
}

func Close() {
	if err := rdb.Close(); err != nil {
		zap.L().Fatal("fail to close redis ", zap.Error(err))
	}
}
