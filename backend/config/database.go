package config

import (
	"dating-question/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        GetEnv("DB_HOST", "localhost"),
        GetEnv("DB_USER", "user"),
        GetEnv("DB_PASSWORD", "password"),
        GetEnv("DB_NAME", "dating_question"),
        GetEnv("DB_PORT", "54321"),
    )
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic(fmt.Sprintf("Failed to connect to database! Error: %v", err))
    }
	database.AutoMigrate(&models.Category{}, &models.Question{}, &models.Session{})
	DB = database
}
