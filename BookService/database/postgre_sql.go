package database

import (
	"book_service/configs"
	"book_service/database/migrations"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func InitiationDatabase(configs configs.ConfigurationsInterface) *gorm.DB {
	// Set dsn (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		configs.DatabaseConfiguration().Host,
		configs.DatabaseConfiguration().Port,
		configs.DatabaseConfiguration().Username,
		configs.DatabaseConfiguration().Database,
		configs.DatabaseConfiguration().Password,
	)

	// RunMigrationCreateBook
	err := migrations.RunMigrationCreateBook(dsn)
	if err != nil {
		log.Fatalf("Migration create book failed: %v", err)
	}
	// RunMigrationCreateBorrowing
	err2 := migrations.RunMigrationCreateBorrowing(dsn)
	if err2 != nil {
		log.Fatalf("Migration create borrowing failed: %v", err2)
	}

	// set gorm
	db, errGorm := gorm.Open("postgres", dsn)
	if errGorm != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	// Return db
	return db
}
