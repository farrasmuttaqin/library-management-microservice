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

type BorrowingHandler struct {
	Ctx              context.Context
	RedisHelper      database.RedisHelper
	BorrowingUseCase usecase.BorrowingUseCase
	Configs          configs.ConfigurationsInterface
}

// BorrowBook handles POST requests to borrow a book
func (h *BorrowingHandler) BorrowBook(c *fiber.Ctx) error {
	var borrowingData models.BorrowingForm

	// Parse and validate the request body
	if err := c.BodyParser(&borrowingData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "message": "Invalid request body"})
	}

	// Borrow the book
	response, err3 := h.BorrowingUseCase.BorrowBook(uint(borrowingData.BookId), uint(borrowingData.UserId))
	if err3 != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// ReturnBook handles POST requests to return a borrowed book
func (h *BorrowingHandler) ReturnBook(c *fiber.Ctx) error {
	borrowingID, err := strconv.ParseUint(c.Params("borrowing_id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid borrowing ID"})
	}

	// Return the book
	response, err := h.BorrowingUseCase.ReturnBook(uint(borrowingID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response)
	}

	return c.Status(http.StatusOK).JSON(response)
}
