package repository

import (
	"backend_scrolling_simulator/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  os.Getenv("SCROLLING_SIMULATOR_API_DB_DSN"),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}))
		if err != nil {
			panic("failed to connect database")
		}

		err = db.AutoMigrate(&models.Like{})
		if err != nil {
			return nil
		}

		err = db.AutoMigrate(&models.CrossPost{})
		if err != nil {
			return nil
		}
	}

	return db
}
