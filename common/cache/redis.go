package cache

import (
	redis "github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	return newRedisClient("localhost:6379", "", 0)
}

func newRedisClient(addr string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return rdb
}
