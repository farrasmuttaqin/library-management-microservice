package route

import (
	"category_service/configs"
	"category_service/helpers/database"
	"category_service/middleware"
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
	// register category router
	RegisterCategoryRouter := CategoryRouter{
		F:           h.F,
		Ctx:         h.Ctx,
		ConfigData:  h.Configs,
		RedisHelper: h.RedisHelper,
		Database:    h.Database,
	}

	RegisterCategoryRouter.CategoryAPIHandler()
	// end register category router

	return h
}

// RegisterAPIMiddleware ...
func (h *HTTPHandler) RegisterAPIMiddleware() {
	h.F.Use(middleware.AccessLogger())
	h.F.Use(middleware.ErrorHandler())
}
