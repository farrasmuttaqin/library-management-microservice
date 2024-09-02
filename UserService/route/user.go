package route

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	configs2 "user_service/configs"
	"user_service/handlers"
	"user_service/helpers/database"
	"user_service/middleware"
	repo "user_service/repositories"
	"user_service/usecase"
)

// UserRouter ...
type UserRouter struct {
	F           *fiber.App
	Ctx         context.Context
	Configs     configs2.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// UserAPIHandler ...
func (router *UserRouter) UserAPIHandler() *UserRouter {
	userRepo := repo.UserRepository{DB: router.Database}
	userUseCase := usecase.UserUseCase{
		Configs:     router.Configs,
		RedisHelper: router.RedisHelper,
		UserRepo:    userRepo,
	}
	userHandler := handlers.UserHandler{
		Ctx:         router.Ctx,
		RedisHelper: router.RedisHelper,
		UserRepo:    userRepo,
		UserUseCase: userUseCase,
		Configs:     router.Configs,
	}

	// Group for user routes
	group := router.F.Group("/api/user")

	// Register endpoint
	group.Post("/register", userHandler.Register)

	// Login endpoint
	group.Post("/login", userHandler.Login)

	// GetUserIdByAuth endpoint
	group.Get("/get", middleware.ValidateJWT(userUseCase), userHandler.GetUserIdByAuth)

	// Validate token endpoint
	group.Post("/validate-token", userHandler.ValidateToken)

	return router
}
