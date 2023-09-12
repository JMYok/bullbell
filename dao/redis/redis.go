package redis

import (
	"bluebell/settings"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		// use default DB
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("redis connect failed", zap.Error(err))
		return err
	}
	return nil
}

func Close() {
	_ = client.Close()
}
