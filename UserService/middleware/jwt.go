package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
	"user_service/helpers"
	"user_service/usecase"
)

// ValidateJWT is a middleware function that validates JWT tokens using a gRPC service.
func ValidateJWT(userUseCase usecase.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from the Authorization header
		token := c.Get("Authorization")
		if helpers.IsEmptyStruct(token) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": 0, "error": "Missing token"})
		}

		// remove Bearer prefix
		token = strings.TrimPrefix(token, "Bearer ")

		// Call ValidateToken gRPC method
		res, err2 := userUseCase.ValidateToken(token)
		if err2 != nil {
			log.Printf("Failed to validate token: %v", err2)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": err2.Error()})
		}

		// Set user ID in the context
		c.Locals("userID", res.UserID)

		// Proceed to the next handler
		return c.Next()
	}
}
