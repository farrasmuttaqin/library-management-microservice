package repo

import (
	"book_service/models"
	"github.com/jinzhu/gorm"
)

type BorrowingRepository struct {
	DB *gorm.DB
}

// tableBorrowing retrieves the table name for the Author model.
func (r *BorrowingRepository) tableBorrowing() string {
	return models.Borrowing{}.TableName()
}

func (r *BorrowingRepository) Create(borrowing *models.Borrowing) error {
	return r.DB.Table(r.tableBorrowing()).Create(borrowing).Error
}

func (r *BorrowingRepository) Get(query map[string]interface{}) (*models.Borrowing, error) {
	var borrowing models.Borrowing
	err := r.DB.Table(r.tableBorrowing()).
		Preload("Book").
		Preload("Book.Author").
		Preload("Book.Category").
		Where(query).First(&borrowing).Error
	if err != nil {
		return nil, err
	}
	return &borrowing, nil
}

func (r *BorrowingRepository) Update(id uint, borrowing *models.Borrowing) error {
	return r.DB.Table(r.tableBorrowing()).Where("id = ?", id).Save(borrowing).Error
}

func (r *BorrowingRepository) Delete(id uint) error {
	return r.DB.Table(r.tableBorrowing()).Where("id = ?", id).Delete(&models.Borrowing{}).Error
}
