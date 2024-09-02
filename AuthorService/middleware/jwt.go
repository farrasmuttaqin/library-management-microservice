package middleware

import (
	"author_service/helpers"
	"context"
	"log"
	"net/http"
	"strings"

	"author_service/configs"
	pb "author_service/proto/user"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

// ValidateJWT is a middleware function that validates JWT tokens using a gRPC service.
func ValidateJWT(config configs.ConfigurationsInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a gRPC connection
		address := config.GRPCConfiguration().UserServiceGRPCAddress
		if address == "" {
			log.Printf("gRPC address is not configured properly")
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": "Service address not configured"})
		}

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Printf("Failed to connect to gRPC service at %s: %v", address, err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": "Service unavailable"})
		}
		defer conn.Close()

		// Create a gRPC client
		client := pb.NewUserServiceClient(conn) // Corrected client creation

		// Extract token from the Authorization header
		token := c.Get("Authorization")
		if helpers.IsEmptyStruct(token) {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": 0, "error": "Missing token"})
		}

		// remove Bearer prefix
		token = strings.TrimPrefix(token, "Bearer ")

		// Call ValidateToken gRPC method
		res, err2 := client.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token})
		if err2 != nil {
			log.Printf("Failed to validate token: %v", err2)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": err2.Error()})
		}
		if res.Status == 0 {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 0, "error": res.ErrorMessage})
		}

		// Set user ID in the context
		c.Locals("userID", res.Claims.UserId)

		// Proceed to the next handler
		return c.Next()
	}
}
