package database

import (
	"fmt"
	"log"
	"user_service/configs"
	"user_service/database/migrations"

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

	db, err2 := gorm.Open("postgres", dsn)
	if err2 != nil {
		log.Fatal("failed to connect to the database:", err2)
	}

	// Return db
	return db
}
