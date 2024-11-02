package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/dddong3/Bid_Backend/config"
	"github.com/dddong3/Bid_Backend/logger"

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
			dbPath := config.GetEnv("SQLITE_DB_PATH", "test.sqlite3")
			db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		default:
			logger.Logger.Fatalf("Unsupported database type: %s", dbType)
		}

		if err != nil {
			logger.Logger.Fatalf("Failed to connect to database: %v", err)
			panic(err)
		}
	})

	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		logger.Logger.Warn("Database is not initialized. Initializing database...")
		return InitDB()
	}
	return db
}
