package route

import (
	"category_service/configs"
	"category_service/handlers"
	"category_service/helpers/database"
	"category_service/middleware"
	repo "category_service/repositories"
	"category_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// CategoryRouter ...
type CategoryRouter struct {
	F           *fiber.App
	Ctx         context.Context
	ConfigData  configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// CategoryAPIHandler ...
func (router *CategoryRouter) CategoryAPIHandler() *CategoryRouter {
	categoryRepo := repo.CategoryRepository{DB: router.Database}
	categoryUseCase := usecase.CategoryUseCase{
		Configs:      router.ConfigData,
		RedisHelper:  router.RedisHelper,
		CategoryRepo: categoryRepo,
	}
	categoryHandler := handlers.CategoryHandler{
		Ctx:             router.Ctx,
		RedisHelper:     router.RedisHelper,
		CategoryUseCase: categoryUseCase,
		Configs:         router.ConfigData,
	}
	// Group for category routes
	group := router.F.Group("/api/category")

	// Category Get by ID
	group.Get("/detail/:id", middleware.ValidateJWT(router.ConfigData), categoryHandler.GetCategoryById)

	// Category Remove
	group.Delete("/remove/:id", middleware.ValidateJWT(router.ConfigData), categoryHandler.DeleteCategory)

	// Category Create
	group.Post("/create", middleware.ValidateJWT(router.ConfigData), categoryHandler.CreateCategory)

	// Category Update
	group.Put("/update/:id", middleware.ValidateJWT(router.ConfigData), categoryHandler.UpdateCategory)

	return router
}
