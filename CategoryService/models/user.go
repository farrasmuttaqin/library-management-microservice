package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"size:255;unique;not null" json:"username"`
	Email     string `gorm:"size:255;unique;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"-"` // Store hashed password
	FirstName string `gorm:"size:255" json:"firstName"`
	LastName  string `gorm:"size:255" json:"lastName"`
	Role      string `gorm:"size:50" json:"role"` // For roles like admin, user, etc.
	CreatedAt string `gorm:"index" json:"createdAt"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `gorm:"index" json:"deletedAt"`
}
