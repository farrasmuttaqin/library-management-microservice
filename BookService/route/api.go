package route

import (
	"book_service/configs"
	"book_service/helpers/database"
	"book_service/middleware"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type HTTPHandler struct {
	F           *fiber.App
	Ctx         context.Context
	Configs     configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	Database    *gorm.DB
}

// RegisterAPIHandler ...
func (h *HTTPHandler) RegisterAPIHandler() *HTTPHandler {
	// register book router
	RegisterBookRouter := BookRouter{
		F:           h.F,
		Ctx:         h.Ctx,
		ConfigData:  h.Configs,
		RedisHelper: h.RedisHelper,
		Database:    h.Database,
	}

	RegisterBookRouter.BookAPIHandler()
	// end register book router

	// register borrowing router
	RegisterBorrowingRouter := BorrowingRouter{
		F:           h.F,
		Ctx:         h.Ctx,
		ConfigData:  h.Configs,
		RedisHelper: h.RedisHelper,
		Database:    h.Database,
	}

	RegisterBorrowingRouter.BorrowingAPIHandler()
	// end register book router

	return h
}

// RegisterAPIMiddleware ...
func (h *HTTPHandler) RegisterAPIMiddleware() {
	h.F.Use(middleware.AccessLogger())
	h.F.Use(middleware.ErrorHandler())
}
