package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/viant/toolbox"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
	"user_service/configs"
	"user_service/database"
	"user_service/helpers"
	database2 "user_service/helpers/database"
	"user_service/helpers/format"
	"user_service/proto/user"
	repo "user_service/repositories"
	"user_service/route"
	"user_service/usecase"
)

// startGRPCServer sets up and starts the gRPC server with reflection
func startGRPCServer(
	config configs.ConfigurationsInterface,
	database *gorm.DB,
	redisHelper database2.RedisHelper) {
	grpcPort := config.GRPCConfiguration().UserServiceGRPCPort
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()

	userRepo := repo.UserRepository{DB: database}
	userUseCase := usecase.UserUseCase{
		Configs:     config,
		RedisHelper: redisHelper,
		UserRepo:    userRepo,
	}

	// Create an instance of UserServiceServer with the UserUseCase
	userServiceServer := &usecase.UserServiceServer{
		UserUseCase: &userUseCase,
	}

	// Register your gRPC services here
	user.RegisterUserServiceServer(grpcServer, userServiceServer)

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	fmt.Printf("gRPC server listening on port %d\n", grpcPort)

	if err2 := grpcServer.Serve(lis); err2 != nil {
		log.Fatalf("failed to serve gRPC server: %v", err2)
	}
}

func main() {
	// starting engine
	fmt.Println("Starting user service engine ...")

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

	// Start the servers in goroutines
	go startGRPCServer(
		configuration,
		gormDatabaseDB,
		redisHelperMain,
	)

	// Start the Fiber app
	portApps := fmt.Sprintf(":%d", configuration.ReadApplicationConfiguration().Port)
	log.Fatal(app.Listen(portApps))
}
