package core

import (
	"context"

	"gvb/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(config config.Redis) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Password,
		DB:       config.DB,
	})

	// 测试连通性
	if err := cache.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return cache
}
