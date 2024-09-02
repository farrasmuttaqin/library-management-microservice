package handlers

import (
	"author_service/configs"
	"author_service/helpers/database"
	"author_service/models"
	"author_service/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	Ctx           context.Context
	RedisHelper   database.RedisHelper
	AuthorUseCase usecase.AuthorUseCase
	Configs       configs.ConfigurationsInterface
}

func (h *AuthorHandler) CreateAuthor(c *fiber.Ctx) error {
	var author models.Author // Define a struct for request body

	// Parse and validate the request body
	if err := c.BodyParser(&author); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Create author
	createdAuthor, err := h.AuthorUseCase.CreateAuthor(author)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(createdAuthor)
	}

	return c.Status(http.StatusCreated).JSON(createdAuthor)
}

func (h *AuthorHandler) GetAuthorById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// get author by id
	authorById, errGetAuthorById := h.AuthorUseCase.GetAuthorById(uint(id))
	if errGetAuthorById != nil {
		return c.Status(http.StatusBadRequest).JSON(authorById)
	}

	return c.Status(http.StatusOK).JSON(authorById)
}

func (h *AuthorHandler) UpdateAuthor(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	var updateData models.Author // Define a struct for update data

	// Parse and validate the request body
	if err2 := c.BodyParser(&updateData); err2 != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid request body"})
	}

	// Update author
	updatedAuthor, err2 := h.AuthorUseCase.UpdateAuthor(uint(id), updateData)
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(updatedAuthor)
	}

	return c.Status(http.StatusOK).JSON(updatedAuthor)
}

func (h *AuthorHandler) DeleteAuthor(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 0, "error": "Invalid ID"})
	}

	// delete author
	deleteAuthor, err2 := h.AuthorUseCase.DeleteAuthor(uint(id))
	if err2 != nil {
		return c.Status(http.StatusInternalServerError).JSON(deleteAuthor)
	}

	return c.Status(http.StatusOK).JSON(deleteAuthor)
}
