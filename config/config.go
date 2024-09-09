package config

import (
	"github.com/dddong3/Bid_Backend/logger"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetLogLevel() string {
	return strings.ToUpper(GetEnv("LOG_LEVEL", "INFO"))
}
