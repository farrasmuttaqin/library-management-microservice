package repo

import (
	"github.com/jinzhu/gorm"
	"user_service/models"
)

type UserRepository struct {
	DB *gorm.DB
}

// tableUser retrieves the table name for the Author model.
func (r *UserRepository) tableUser() string {
	return models.User{}.TableName()
}

func (r *UserRepository) Get(query map[string]interface{}) (*models.User, error) {
	var user models.User
	err := r.DB.Table(r.tableUser()).Where(query).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
