package migrations

import (
	"author_service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Migrate performs the migration.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Author{})
}

// RunMigration runs the migration using the PostgreSQL database connection.
func RunMigration(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Perform the migration
	return Migrate(db)
}
