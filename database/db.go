package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/dddong3/Bid_Backend/config"
	"github.com/dddong3/Bid_Backend/logger"
	"github.com/dddong3/Bid_Backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dbType := config.GetEnv("DB_TYPE", "sqlite")
		var err error

		switch dbType {
		case "postgres":
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_PORT"),
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		case "sqlite":
			dbPath := os.Getenv("SQLITE_DB_PATH")
			db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		}

		if err != nil {
			logger.Logger.Fatalf("Failed to connect to database: %v", err)
		}
		db.AutoMigrate(&models.AuctionItem{})
		db.AutoMigrate(&models.AuctionItemAnnouncementFile{})
	})

	return db
}

func GetDB() *gorm.DB {
	return db
}
