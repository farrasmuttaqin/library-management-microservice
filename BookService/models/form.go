package models

type BorrowingForm struct {
	UserId int `json:"user_id" validate:"required"`
	BookId int `json:"book_id" validate:"required"`
}
