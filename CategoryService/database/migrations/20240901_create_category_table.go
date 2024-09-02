package migrations

import (
	"category_service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MigrateCreateCategory performs the migration.
func MigrateCreateCategory(db *gorm.DB) error {
	return db.AutoMigrate(&models.Category{})
}

// RunMigrationCreateCategory runs the migration using the PostgreSQL database connection.
func RunMigrationCreateCategory(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Perform the migration
	return MigrateCreateCategory(db)
}
