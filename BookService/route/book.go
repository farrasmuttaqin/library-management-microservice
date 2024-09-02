package route

import (
	"book_service/configs"
	"book_service/handlers"
	"book_service/helpers/database"
	"book_service/middleware"
	repo "book_service/repositories"
	"book_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// BookRouter ...
type BookRouter struct {
	F           *fiber.App
	Ctx         context.Context
	ConfigData  configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// BookAPIHandler ...
func (router *BookRouter) BookAPIHandler() *BookRouter {
	bookRepo := repo.BookRepository{DB: router.Database}
	bookUseCase := usecase.BookUseCase{
		Configs:     router.ConfigData,
		RedisHelper: router.RedisHelper,
		BookRepo:  bookRepo,
	}
	bookHandler := handlers.BookHandler{
		Ctx:           router.Ctx,
		RedisHelper:   router.RedisHelper,
		BookUseCase: bookUseCase,
		Configs:       router.ConfigData,
	}
	// Group for book routes
	group := router.F.Group("/api/book")

	// Book Get by ID
	group.Get("/detail/:id", middleware.ValidateJWT(router.ConfigData), bookHandler.GetBookById)

	// Book Remove
	group.Delete("/remove/:id", middleware.ValidateJWT(router.ConfigData), bookHandler.DeleteBook)

	// Book Create
	group.Post("/create", middleware.ValidateJWT(router.ConfigData), bookHandler.CreateBook)

	// Book Update
	group.Put("/update/:id", middleware.ValidateJWT(router.ConfigData), bookHandler.UpdateBook)

	return router
}
