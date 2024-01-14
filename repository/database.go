package repository

import (
	"backend_scrolling_simulator/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB = nil

var gormLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,        // Don't include params in the SQL log
		Colorful:                  false,       // Disable color
	},
)

func GetDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  os.Getenv("SCROLLING_SIMULATOR_API_DB_DSN"),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), /*, &gorm.Config{
			Logger: gormLogger,
		}*/)
		if err != nil {
			panic("failed to connect database")
		}

		err = db.AutoMigrate(&models.Like{},
			&models.CrossPost{},
			&models.User{},
			&models.Achievement{},
			&models.LeaderboardEntry{},
			&models.Post{},
			&models.Tag{},
			&models.Topic{},
			&models.ReactText{},
		)
		if err != nil {
			return nil
		}
	}

	return db
}
