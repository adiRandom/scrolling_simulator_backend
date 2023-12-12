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
		dsn := os.Getenv("SCROLLING_SIMULATOR_API_DB_DSN")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
