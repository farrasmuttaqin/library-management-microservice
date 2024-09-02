package models

import (
	"time"
)

// Borrowing represents the borrowing transaction model for the database.
type Borrowing struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	BookID     uint       `gorm:"not null" json:"book_id"`
	UserID     uint       `gorm:"not null" json:"user_id"`
	BorrowedAt time.Time  `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	Book Book `gorm:"foreignKey:BookID;references:ID" json:"book"`
}

// TableName sets the table name for the Borrowing model.
func (Borrowing) TableName() string {
	return "borrowing"
}
