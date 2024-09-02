package migrations

import (
	"book_service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateCreateBook performs the migration.
func MigrateCreateBook(db *gorm.DB) error {
	return db.AutoMigrate(&models.Book{})
}

// RunMigrationCreateBook runs the migration using the PostgreSQL database connection.
func RunMigrationCreateBook(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Perform the migration
	return MigrateCreateBook(db)
}
