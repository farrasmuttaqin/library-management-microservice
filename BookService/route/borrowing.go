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

// BorrowingRouter ...
type BorrowingRouter struct {
	F           *fiber.App
	Ctx         context.Context
	ConfigData  configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// BorrowingAPIHandler ...
func (router *BorrowingRouter) BorrowingAPIHandler() *BorrowingRouter {
	borrowingRepo := repo.BorrowingRepository{DB: router.Database}
	userRepo := repo.UserRepo{Config: router.ConfigData, Ctx: router.Ctx}
	bookRepo := repo.BookRepository{DB: router.Database}
	borrowingUseCase := usecase.BorrowingUseCase{
		Configs:       router.ConfigData,
		RedisHelper:   router.RedisHelper,
		BorrowingRepo: borrowingRepo,
		UserRepo:      userRepo,
		BookRepo:      bookRepo,
	}
	borrowingHandler := handlers.BorrowingHandler{
		Ctx:              router.Ctx,
		RedisHelper:      router.RedisHelper,
		BorrowingUseCase: borrowingUseCase,
		Configs:          router.ConfigData,
	}
	// Group for book routes
	group := router.F.Group("/api/borrowing")

	// Borrowing routes
	group.Post("/borrow", middleware.ValidateJWT(router.ConfigData), borrowingHandler.BorrowBook) // Route to borrow a book
	group.Post("/return/:borrowing_id", middleware.ValidateJWT(router.ConfigData), borrowingHandler.ReturnBook)

	return router
}
