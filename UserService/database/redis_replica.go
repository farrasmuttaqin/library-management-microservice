package database

import (
	"fmt"
	"user_service/configs"

	"github.com/redis/go-redis/v9"
)

func InitiationRedisReplica(configs configs.ConfigurationsInterface) *redis.Client {
	// set address
	address := fmt.Sprintf("%s:%d",
		configs.RedisReplicaConfiguration().Host,
		configs.RedisReplicaConfiguration().Port,
	)

	// set password
	password := configs.RedisReplicaConfiguration().Password

	// set client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return redisClient
}
