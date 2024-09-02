package route

import (
	"author_service/configs"
	"author_service/helpers/database"
	"author_service/middleware"
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
	// register author router
	RegisterAuthorRouter := AuthorRouter{
		F:           h.F,
		Ctx:         h.Ctx,
		ConfigData:  h.Configs,
		RedisHelper: h.RedisHelper,
		Database:    h.Database,
	}

	RegisterAuthorRouter.AuthorAPIHandler()
	// end register author router

	return h
}

// RegisterAPIMiddleware ...
func (h *HTTPHandler) RegisterAPIMiddleware() {
	h.F.Use(middleware.AccessLogger())
	h.F.Use(middleware.ErrorHandler())
}
