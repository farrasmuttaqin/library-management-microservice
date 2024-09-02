package migrations

import (
	"book_service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateCreateBorrowing performs the migration.
func MigrateCreateBorrowing(db *gorm.DB) error {
	return db.AutoMigrate(&models.Borrowing{})
}

// RunMigrationCreateBorrowing runs the migration using the PostgreSQL database connection.
func RunMigrationCreateBorrowing(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Perform the migration
	return MigrateCreateBorrowing(db)
}
