package database

import (
	"author_service/configs"
	"author_service/database/migrations"
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

	// run migration
	err := migrations.RunMigration(dsn)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// set gorm
	db, errGorm := gorm.Open("postgres", dsn)
	if errGorm != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	// Return db
	return db
}
