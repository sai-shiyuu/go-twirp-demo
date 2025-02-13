package cache

import (
	"fmt"
	"go-web/common/config"

	redis "github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	return newRedisClient(fmt.Sprintf("%s:%d",
		config.Configs.Redis.Host,
		config.Configs.Redis.Port),
		config.Configs.Redis.Password,
		config.Configs.Redis.DB)
}

func newRedisClient(addr string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return rdb
}
