package usecase

import (
	"book_service/configs"
	"book_service/helpers"
	"book_service/helpers/database"
	"book_service/models"
	repo "book_service/repositories"
	"errors"
	"time"
)

type BorrowingUseCase struct {
	Configs       configs.ConfigurationsInterface
	RedisHelper   database.RedisHelper
	BookRepo      repo.BookRepository
	UserRepo      repo.UserRepo
	BorrowingRepo repo.BorrowingRepository
}

func (u *BorrowingUseCase) BorrowBook(bookID uint, userID uint) (map[string]interface{}, error) {
	// Check if the book exists
	bookQuery := map[string]interface{}{
		"id": bookID,
	}
	book, err := u.BookRepo.Get(bookQuery)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	if helpers.IsEmptyStruct(book) {
		return map[string]interface{}{
			"status":  0,
			"message": "Book not found",
		}, errors.New("book not found")
	}

	// Check if the book is available
	if book.StockQuantity <= 0 {
		return map[string]interface{}{
			"status":  0,
			"message": "Book is out of stock",
		}, errors.New("book is out of stock")
	}

	user, err2 := u.UserRepo.Get(int(userID))
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err2.Error(),
		}, err2
	}

	if helpers.IsEmptyStruct(user) {
		return map[string]interface{}{
			"status":  0,
			"message": "User not found",
		}, errors.New("user not found")
	}

	// Create borrowing record
	borrowing := models.Borrowing{
		BookID:     bookID,
		UserID:     userID,
		BorrowedAt: time.Now(),
	}

	if err3 := u.BorrowingRepo.Create(&borrowing); err3 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err3.Error(),
		}, err3
	}

	book.StockQuantity--
	// Update book stock
	update := map[string]interface{}{
		"stock_quantity": book.StockQuantity,
	}
	if err3 := u.BookRepo.UpdateSelective(bookID, book, update); err3 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err3.Error(),
		}, err3
	}

	// get again borrowing data
	q := map[string]interface{}{
		"id": borrowing.ID,
	}
	borrowingData, _ := u.BorrowingRepo.Get(q)

	return map[string]interface{}{
		"status":    1,
		"message":   "Book borrowed successfully",
		"borrowing": borrowingData,
		"user":      user,
	}, nil
}

func (u *BorrowingUseCase) ReturnBook(borrowingID uint) (map[string]interface{}, error) {
	// Check if the borrowing record exists
	borrowingQuery := map[string]interface{}{
		"id": borrowingID,
	}
	borrowing, err := u.BorrowingRepo.Get(borrowingQuery)
	if err != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err.Error(),
		}, err
	}

	if helpers.IsEmptyStruct(borrowing) {
		return map[string]interface{}{
			"status":  0,
			"message": "Borrowing record not found",
		}, errors.New("borrowing record not found")
	}

	if borrowing.ReturnedAt != nil {
		return map[string]interface{}{
			"status":  0,
			"message": "Book already returned",
		}, errors.New("book already returned")
	}

	// Update return date
	now := time.Now()
	borrowing.ReturnedAt = &now
	if err2 := u.BorrowingRepo.Update(borrowingID, borrowing); err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err2.Error(),
		}, err2
	}

	// Update book stock
	bookQuery := map[string]interface{}{
		"id": borrowing.BookID,
	}
	book, err2 := u.BookRepo.Get(bookQuery)
	if err2 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err2.Error(),
		}, err2
	}

	if helpers.IsEmptyStruct(book) {
		return map[string]interface{}{
			"status":  0,
			"message": "Book not found",
		}, errors.New("book not found")
	}

	book.StockQuantity++
	// Update book stock
	update := map[string]interface{}{
		"stock_quantity": book.StockQuantity,
	}
	if err3 := u.BookRepo.UpdateSelective(book.ID, book, update); err3 != nil {
		return map[string]interface{}{
			"status":  0,
			"message": err3.Error(),
		}, err3
	}

	user, _ := u.UserRepo.Get(int(borrowing.UserID))
	// get again borrowing data
	q := map[string]interface{}{
		"id": borrowing.ID,
	}
	borrowingData, _ := u.BorrowingRepo.Get(q)

	return map[string]interface{}{
		"status":    1,
		"message":   "Book returned successfully",
		"borrowing": borrowingData,
		"user":      user,
	}, nil
}
