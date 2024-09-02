package repo

import (
	"author_service/models"
	"github.com/jinzhu/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

// tableAuthor retrieves the table name for the Author model.
func (r *AuthorRepository) tableAuthor() string {
	return models.Author{}.TableName()
}

func (r *AuthorRepository) GetExceptId(id uint, query map[string]interface{}) (*models.Author, error) {
	var author models.Author
	err := r.DB.Table(r.tableAuthor()).Where(query).Where("id != ?", id).First(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) Create(author *models.Author) error {
	return r.DB.Table(r.tableAuthor()).Create(author).Error
}

func (r *AuthorRepository) Get(query map[string]interface{}) (*models.Author, error) {
	var author models.Author
	err := r.DB.Table(r.tableAuthor()).Where(query).First(&author).Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) Update(id uint, author *models.Author) error {
	resultUpdate := r.DB.Table(r.tableAuthor()).Where("id = ?", id).Updates(author).Error
	_ = r.DB.Table(r.tableAuthor()).Where("id = ?", id).First(&author).Error
	return resultUpdate
}

func (r *AuthorRepository) Delete(id uint) error {
	return r.DB.Table(r.tableAuthor()).Where("id = ?", id).Delete(&models.Author{}).Error
}
