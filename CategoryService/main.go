package main

import (
	"category_service/configs"
	"category_service/database"
	"category_service/helpers"
	database2 "category_service/helpers/database"
	"category_service/helpers/format"
	"category_service/route"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/viant/toolbox"
	"log"
	"os"
	"time"
)

func main() {
	// starting engine
	fmt.Println("Starting book service engine ...")

	// set context
	ctx := context.Background()

	// set fiber app
	app := fiber.New()

	// Load Configuration
	configuration := configs.NewConfigViper(viper.New())
	redisMasterClient := database.InitiationRedisMaster(configuration)
	redisReplicaClient := database.InitiationRedisReplica(configuration)
	gormDatabaseDB := database.InitiationDatabase(configuration)

	// set timezone
	// if empty, change timezone to jakarta
	var timeZoneApp string
	if helpers.IsEmptyStruct(configuration.ReadApplicationConfiguration().TimeZone) {
		timeZoneApp = "Asia/Jakarta"
	} else {
		timeZoneApp = configuration.ReadApplicationConfiguration().TimeZone
	}
	location, err := time.LoadLocation(timeZoneApp)
	if err != nil {
		log.Fatalf("Failed to load time zone : %v", err)
	}
	time.Local = location

	// set logs path
	logTimeLayout := toolbox.DateFormatToLayout(format.DateFormat)
	currentDate := time.Now().Format(logTimeLayout)
	logFile, _ := os.OpenFile(format.PublicLogPath+currentDate+format.LogExtension, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// set redis helper
	redisHelperMain := database2.RedisHelper{
		RedisMaster:  redisMasterClient,
		RedisReplica: redisReplicaClient,
		Ctx:          ctx,
	}

	// set main router
	mainRouter := route.HTTPHandler{
		F:           app,
		Ctx:         ctx,
		Configs:     configuration,
		RedisHelper: redisHelperMain,
		Database:    gormDatabaseDB,
	}

	// Register middleware and routes
	mainRouter.RegisterAPIMiddleware()
	mainRouter.RegisterAPIHandler()

	// Start the Fiber app
	portApps := fmt.Sprintf(":%d", configuration.ReadApplicationConfiguration().Port)
	log.Fatal(app.Listen(portApps))
}
