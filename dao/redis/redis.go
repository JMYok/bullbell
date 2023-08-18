package redis

import (
	"bluebell/settings"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port),
		Password: settings.Conf.RedisConfig.Password,
		// use default DB
		DB:       settings.Conf.RedisConfig.DB,
		PoolSize: settings.Conf.RedisConfig.PoolSize,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
