package repo

import (
	"category_service/models"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

// tableCategory retrieves the table name for the Author model.
func (r *CategoryRepository) tableCategory() string {
	return models.Category{}.TableName()
}

func (r *CategoryRepository) Create(book *models.Category) error {
	return r.DB.Table(r.tableCategory()).Create(book).Error
}

func (r *CategoryRepository) Get(query map[string]interface{}) (*models.Category, error) {
	var category models.Category
	err := r.DB.Table(r.tableCategory()).
		Where(query).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetExceptId(id uint, query map[string]interface{}) (*models.Category, error) {
	var category models.Category
	err := r.DB.Table(r.tableCategory()).Where(query).Where("id != ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) Update(id uint, book *models.Category) error {
	resultUpdate := r.DB.Table(r.tableCategory()).Where("id = ?", id).Updates(book).Error
	_ = r.DB.Table(r.tableCategory()).Where("id = ?", id).First(&book).Error
	return resultUpdate
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Table(r.tableCategory()).Where("id = ?", id).Delete(&models.Category{}).Error
}
