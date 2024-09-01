package database

import (
	"fmt"
	"log"
	"sa_telco_legacy/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitiationDatabase(configs configs.ConfigurationsInterface) *gorm.DB {
	// set dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.DatabaseConfiguration().Username,
		configs.DatabaseConfiguration().Password,
		configs.DatabaseConfiguration().Host,
		configs.DatabaseConfiguration().Port,
		configs.DatabaseConfiguration().Database,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to the database.")
	}
	// return db
	return db
}
