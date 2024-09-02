package route

import (
	"author_service/configs"
	"author_service/handlers"
	"author_service/helpers/database"
	"author_service/middleware"
	repo "author_service/repositories"
	"author_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// AuthorRouter ...
type AuthorRouter struct {
	F           *fiber.App
	Ctx         context.Context
	ConfigData  configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// AuthorAPIHandler ...
func (router *AuthorRouter) AuthorAPIHandler() *AuthorRouter {
	authorRepo := repo.AuthorRepository{DB: router.Database}
	authorUseCase := usecase.AuthorUseCase{
		Configs:     router.ConfigData,
		RedisHelper: router.RedisHelper,
		AuthorRepo:  authorRepo,
	}
	authorHandler := handlers.AuthorHandler{
		Ctx:           router.Ctx,
		RedisHelper:   router.RedisHelper,
		AuthorUseCase: authorUseCase,
		Configs:       router.ConfigData,
	}
	// Group for author routes
	group := router.F.Group("/api/author")

	// Author Get by ID
	group.Get("/detail/:id", middleware.ValidateJWT(router.ConfigData), authorHandler.GetAuthorById)

	// Author Remove
	group.Delete("/remove/:id", middleware.ValidateJWT(router.ConfigData), authorHandler.DeleteAuthor)

	// Author Create
	group.Post("/create", middleware.ValidateJWT(router.ConfigData), authorHandler.CreateAuthor)

	// Author Update
	group.Put("/update/:id", middleware.ValidateJWT(router.ConfigData), authorHandler.UpdateAuthor)

	return router
}
