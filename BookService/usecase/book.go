package usecase

import (
	"book_service/configs"
	"book_service/helpers"
	"book_service/helpers/database"
	"book_service/models"
	repo "book_service/repositories"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type BookUseCase struct {
	Configs     configs.ConfigurationsInterface
	RedisHelper database.RedisHelper
	BookRepo    repo.BookRepository
}

func (u *BookUseCase) GetBookById(id uint) (map[string]interface{}, error) {
	cacheKey := "book:" + strconv.FormatUint(uint64(id), 10)

	// Try to get the data from Redis cache
	content, err := u.RedisHelper.Get(cacheKey).Result()
	if err == nil {
		var book models.Book
		if err2 := json.Unmarshal([]byte(content), &book); err2 == nil && !helpers.IsEmptyStruct(book) {
			return map[string]interface{}{
				"status":  1,
				"message": "success",
				"book":    book,
			}, nil
		}
	}

	// If not in cache, get from database
	qId := map[string]interface{}{
		"id": id,
	}
	book, err2 := u.BookRepo.Get(qId)
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to retrieve book",
		}, err2
	}

	if helpers.IsEmptyStruct(book) {
		return map[string]interface{}{
			"status":  0,
			"message": "Book not found",
		}, errors.New("book is empty")
	}

	// Cache the data
	authorJSON, _ := json.Marshal(book)
	cacheTTL := time.Duration(u.Configs.RedisMasterConfiguration().TTLSecond) * time.Second
	if err3 := u.RedisHelper.Set(cacheKey, string(authorJSON), cacheTTL); err3 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "success get data, but failed to cache data",
			"book":    book,
		}, err3
	}

	return map[string]interface{}{
		"status":  1,
		"message": "success",
		"author":  book,
	}, nil
}

func (u *BookUseCase) CreateBook(bookData models.Book) (map[string]interface{}, error) {
	// Check if the author with the same isbn already exists
	qValidation := map[string]interface{}{
		"isbn": bookData.ISBN,
	}
	existingBookISBN, _ := u.BookRepo.Get(qValidation)
	if !helpers.IsEmptyStruct(existingBookISBN) {
		return map[string]interface{}{
			"status":  0,
			"message": "book with this isbn already exists",
		}, errors.New("book with this isbn already exists")
	}

	// Create book in the database
	err := u.BookRepo.Create(&bookData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	return map[string]interface{}{
		"status":  1,
		"message": "Book created successfully",
		"book":    bookData,
	}, nil
}

func (u *BookUseCase) UpdateBook(id uint, bookData models.Book) (map[string]interface{}, error) {
	// Check
	qValidation := map[string]interface{}{
		"id": id,
	}
	exists, _ := u.BookRepo.Get(qValidation)
	if helpers.IsEmptyStruct(exists) {
		return map[string]interface{}{
			"status":  0,
			"message": "Book not found",
		}, errors.New("book not found")
	}

	// Check if the author with the same isbn already exists
	qValidation = map[string]interface{}{
		"isbn": bookData.ISBN,
	}
	existingBookISBN, _ := u.BookRepo.GetExceptId(id, qValidation)
	if !helpers.IsEmptyStruct(existingBookISBN) {
		return map[string]interface{}{
			"status":  0,
			"message": "book with this isbn already exists",
		}, errors.New("book with this isbn already exists")
	}

	// Update book in the database
	err := u.BookRepo.Update(id, &bookData)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to update book",
		}, err
	}

	cacheKey := "book:" + strconv.FormatUint(uint64(id), 10)
	_ = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Book updated successfully",
		"author":  bookData,
	}, nil
}

func (u *BookUseCase) DeleteBook(id uint) (map[string]interface{}, error) {
	// Delete book from the database
	if err := u.BookRepo.Delete(id); err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Failed to delete book",
		}, err
	}

	cacheKey := "book:" + strconv.FormatUint(uint64(id), 10)
	_ = u.RedisHelper.Del(cacheKey)

	return map[string]interface{}{
		"status":  1,
		"message": "Book deleted successfully",
	}, nil
}
