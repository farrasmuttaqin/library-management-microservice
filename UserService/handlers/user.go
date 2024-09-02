package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"user_service/configs"
	"user_service/helpers/database"
	"user_service/models"
	repo "user_service/repositories"
	"user_service/usecase"
)

type UserHandler struct {
	Ctx         context.Context
	RedisHelper database.RedisHelper
	UserRepo    repo.UserRepository
	UserUseCase usecase.UserUseCase
	Configs     configs.ConfigurationsInterface
}

// Login handles user login and token generation
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginData models.LoginForm

	// Parse and validate the request body
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "message": "Invalid request body"})
	}

	// Login user and generate token
	response, err := h.UserUseCase.Login(loginData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// GetUserIdByAuth ..
func (h *UserHandler) GetUserIdByAuth(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	// get response
	response, err := h.UserUseCase.GetByUserId(userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"status": 0,
			"error":  err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status": 1,
		"user":   response,
	})
}

// Register handles user registration
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var registerData models.RegisterForm

	// Parse and validate the request body
	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "message": "Invalid request body"})
	}

	// Register user
	response, err := h.UserUseCase.Register(registerData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// ValidateToken handles the token validation.
func (h *UserHandler) ValidateToken(c *fiber.Ctx) error {
	// Define a variable to hold the request body
	var req models.ValidateToken

	// Parse the request body into the TokenRequest struct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "message": "Invalid request body"})
	}

	// Check if the token is provided
	if req.JWTToken == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "message": "jwt_token is required"})
	}

	// Validate the token
	claims, err := h.UserUseCase.ValidateToken(req.JWTToken)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": 0, "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": 1, "claims": claims})
}
