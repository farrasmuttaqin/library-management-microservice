package database

import (
	"book_service/configs"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitiationRedisMaster(configs configs.ConfigurationsInterface) *redis.Client {
	// set address
	address := fmt.Sprintf("%s:%d",
		configs.RedisMasterConfiguration().Host,
		configs.RedisMasterConfiguration().Port,
	)

	// set password
	password := configs.RedisMasterConfiguration().Password

	// set client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return redisClient
}
