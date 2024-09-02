package handlers

import (
	"category_service/configs"
	"category_service/helpers/database"
	"category_service/models"
	"category_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	Ctx             context.Context
	RedisHelper     database.RedisHelper
	CategoryUseCase usecase.CategoryUseCase
	Configs         configs.ConfigurationsInterface
}

// CreateCategory handles POST requests to create a new book
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	// Parse and validate the request body
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Create category
	response, err := h.CategoryUseCase.CreateCategory(category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// GetCategoryById handles GET requests to retrieve a book by its ID
func (h *CategoryHandler) GetCategoryById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// Get category by ID
	response, err2 := h.CategoryUseCase.GetCategoryById(uint(id))
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// UpdateCategory handles PUT requests to update an existing book
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	var categoryData models.Category

	// Parse and validate the request body
	if err2 := c.BodyParser(&categoryData); err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Update category
	response, err3 := h.CategoryUseCase.UpdateCategory(uint(id), categoryData)
	if err3 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// DeleteCategory handles DELETE requests to remove a book by its ID
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// Delete category
	response, err2 := h.CategoryUseCase.DeleteCategory(uint(id))
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}
