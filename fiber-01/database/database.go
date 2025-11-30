package database

import (
	"log"
	"github.com/paung29/model"
	"time"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info,   // Log level
			Colorful:      true,          // Color
		},
	)

	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(
		&model.Account{},
		&model.User{},
		&model.Post{},
	)

	DB = db
}