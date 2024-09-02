package handlers

import (
	"book_service/configs"
	"book_service/helpers/database"
	"book_service/models"
	"book_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type BookHandler struct {
	Ctx         context.Context
	RedisHelper database.RedisHelper
	BookUseCase usecase.BookUseCase
	Configs     configs.ConfigurationsInterface
}

// CreateBook handles POST requests to create a new book
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var book models.Book

	// Parse and validate the request body
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Create book
	response, err := h.BookUseCase.CreateBook(book)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusCreated).JSON(response)
}

// GetBookById handles GET requests to retrieve a book by its ID
func (h *BookHandler) GetBookById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// Get book by ID
	response, err2 := h.BookUseCase.GetBookById(uint(id))
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// UpdateBook handles PUT requests to update an existing book
func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	var bookData models.Book

	// Parse and validate the request body
	if err2 := c.BodyParser(&bookData); err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Update book
	response, err3 := h.BookUseCase.UpdateBook(uint(id), bookData)
	if err3 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// DeleteBook handles DELETE requests to remove a book by its ID
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// Delete book
	response, err2 := h.BookUseCase.DeleteBook(uint(id))
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}
