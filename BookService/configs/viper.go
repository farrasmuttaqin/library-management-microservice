package configs

import (
	"book_service/models"
	"github.com/spf13/viper"
	"log"
)

type ConfigurationsInterface interface {
	ReadApplicationConfiguration() models.ApplicationConfiguration
	DatabaseConfiguration() models.DatabaseMySQLConfiguration
	RedisMasterConfiguration() models.RedisConfiguration
	RedisReplicaConfiguration() models.RedisConfiguration
	GRPCConfiguration() models.GRPCConfiguration
}

type viperConfiguration struct {
	viper *viper.Viper
}

func NewConfigViper(v *viper.Viper) ConfigurationsInterface {
	// viper setup automatic env
	v.AutomaticEnv()

	// set config file name
	configFileName := ".env"

	// set configuration path
	v.SetConfigType("env")
	v.SetConfigName(configFileName)
	v.AddConfigPath("./")

	// viper read .env
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read environment variable")
	}

	return &viperConfiguration{v}
}

func (v *viperConfiguration) ReadApplicationConfiguration() models.ApplicationConfiguration {
	return models.ApplicationConfiguration{
		Name:           v.viper.GetString("APP_NAME"),
		Port:           v.viper.GetInt("APP_PORT"),
		ServerTimeHour: v.viper.GetInt("APP_SERVER_TIME_HOUR"),
		ClientTimeHour: v.viper.GetInt("APP_CLIENT_TIME_HOUR"),
		TimeZone:       v.viper.GetString("TIME_ZONE"),
	}
}

func (v *viperConfiguration) DatabaseConfiguration() models.DatabaseMySQLConfiguration {
	return models.DatabaseMySQLConfiguration{
		Driver:   v.viper.GetString("DB_DRIVER"),
		Host:     v.viper.GetString("DB_HOST"),
		Port:     v.viper.GetInt("DB_PORT"),
		Username: v.viper.GetString("DB_USERNAME"),
		Password: v.viper.GetString("DB_PASSWORD"),
		Database: v.viper.GetString("DB_DATABASE"),
	}
}

func (v *viperConfiguration) RedisMasterConfiguration() models.RedisConfiguration {
	return models.RedisConfiguration{
		Host:      v.viper.GetString("REDIS_MASTER_HOST"),
		Port:      v.viper.GetInt("REDIS_MASTER_PORT"),
		Password:  v.viper.GetString("REDIS_MASTER_PASSWORD"),
		TTLSecond: v.viper.GetInt("REDIS_MASTER_TTL_SECOND"),
	}
}

func (v *viperConfiguration) RedisReplicaConfiguration() models.RedisConfiguration {
	return models.RedisConfiguration{
		Host:      v.viper.GetString("REDIS_REPLICA_HOST"),
		Port:      v.viper.GetInt("REDIS_REPLICA_PORT"),
		Password:  v.viper.GetString("REDIS_REPLICA_PASSWORD"),
		TTLSecond: v.viper.GetInt("REDIS_REPLICA_TTL_SECOND"),
	}
}

func (v *viperConfiguration) GRPCConfiguration() models.GRPCConfiguration {
	return models.GRPCConfiguration{
		UserServiceGRPCAddress: v.viper.GetString("GRPC_USER_SERVICE_GRPC_ADDRESS"),
	}
}
