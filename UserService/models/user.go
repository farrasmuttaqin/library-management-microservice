package models

import "gorm.io/gorm"

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:255;unique;not null" json:"username"`
	Email     string         `gorm:"size:255;unique;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"` // Store hashed password
	FirstName string         `gorm:"size:255" json:"firstName"`
	LastName  string         `gorm:"size:255" json:"lastName"`
	Role      string         `gorm:"size:50" json:"role"` // For roles like admin, user, etc.
	CreatedAt gorm.DeletedAt `gorm:"index" json:"createdAt"`
	UpdatedAt gorm.DeletedAt `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

// TableName sets the table name for the User model.
func (User) TableName() string {
	return "users"
}
