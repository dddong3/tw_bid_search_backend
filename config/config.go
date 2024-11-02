package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func init() {
	LoadEnv()
}

func LoadEnv() {
	if os.Getenv("ENV") == strings.ToLower("production") {
		return
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file")
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
