package repo

import (
	"book_service/models"
	"github.com/jinzhu/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

// tableBook retrieves the table name for the Author model.
func (r *BookRepository) tableBook() string {
	return models.Book{}.TableName()
}

func (r *BookRepository) Create(book *models.Book) error {
	resultCreate := r.DB.Table(r.tableBook()).Create(book).Error
	_ = r.DB.Table(r.tableBook()).Preload("Author").Preload("Category").Where("id = ?", book.ID).First(&book).Error
	return resultCreate
}

func (r *BookRepository) Get(query map[string]interface{}) (*models.Book, error) {
	var book models.Book
	err := r.DB.Table(r.tableBook()).Preload("Author").Preload("Category").
		Where(query).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetExceptId(id uint, query map[string]interface{}) (*models.Book, error) {
	var book models.Book
	err := r.DB.Table(r.tableBook()).Where(query).Where("id != ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(id uint, book *models.Book) error {
	resultUpdate := r.DB.Table(r.tableBook()).Where("id = ?", id).Updates(book).Error
	_ = r.DB.Table(r.tableBook()).Preload("Author").Preload("Category").Where("id = ?", id).First(&book).Error
	return resultUpdate
}

func (r *BookRepository) UpdateSelective(id uint, book *models.Book, update map[string]interface{}) error {
	resultUpdate := r.DB.Table(r.tableBook()).Where("id = ?", id).Updates(update).Error
	_ = r.DB.Table(r.tableBook()).Preload("Author").Preload("Category").Where("id = ?", id).First(&book).Error
	return resultUpdate
}

func (r *BookRepository) Delete(id uint) error {
	return r.DB.Table(r.tableBook()).Where("id = ?", id).Delete(&models.Book{}).Error
}
