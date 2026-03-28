package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment variables")
    }
}
func GetEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}