package migrations

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"user_service/models"
)

// Migrate performs the migration.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
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
