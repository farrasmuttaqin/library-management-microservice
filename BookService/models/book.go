package models

import "time"

// Book represents the book model for the database.
type Book struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	Title         string     `gorm:"size:255;not null" json:"title"`
	ISBN          string     `gorm:"size:20;unique;not null" json:"isbn"`
	AuthorID      uint       `gorm:"not null" json:"author_id"`
	CategoryID    uint       `gorm:"not null" json:"category_id"`
	PublishedDate time.Time  `json:"published_date"`
	Price         float64    `json:"price"`
	StockQuantity int        `gorm:"not null" json:"stock_quantity"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	Author   Author   `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
}

// TableName sets the table name for the Book model.
func (Book) TableName() string {
	return "book"
}
