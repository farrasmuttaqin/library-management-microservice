package route

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"user_service/configs"
	"user_service/helpers/database"
	"user_service/middleware"
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
	// register user router
	RegisterUserRouter := UserRouter{
		F:           h.F,
		Ctx:         h.Ctx,
		Configs:     h.Configs,
		RedisHelper: h.RedisHelper,
		Database:    h.Database,
	}

	RegisterUserRouter.UserAPIHandler()
	// end register user router

	return h
}

// RegisterAPIMiddleware ...
func (h *HTTPHandler) RegisterAPIMiddleware() {
	h.F.Use(middleware.AccessLogger())
	h.F.Use(middleware.ErrorHandler())
}
