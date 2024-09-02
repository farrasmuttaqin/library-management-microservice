package database

import (
	"context"
	"time"
	"user_service/configs"
	"user_service/helpers"

	"github.com/redis/go-redis/v9"
)

type RedisHelper struct {
	RedisMaster  *redis.Client
	RedisReplica *redis.Client
	Ctx          context.Context
	Config       configs.ConfigurationsInterface
}

func (r *RedisHelper) Set(cacheKey string, cacheData interface{}, timeToLeave time.Duration) error {
	if helpers.IsEmptyStruct(timeToLeave) {
		timeToLeave = time.Duration(r.Config.RedisMasterConfiguration().TTLSecond) * time.Second
	}
	return r.RedisMaster.Set(r.Ctx, cacheKey, cacheData, timeToLeave*time.Second).Err()
}

func (r *RedisHelper) Get(cacheKey string) *redis.StringCmd {
	return r.RedisReplica.Get(r.Ctx, cacheKey)
}

func (r *RedisHelper) Del(cacheKey string) error {
	return r.RedisMaster.Del(r.Ctx, cacheKey).Err()
}
