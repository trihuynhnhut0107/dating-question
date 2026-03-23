package main

import (
	"dating-question-backend/internal/config"
	"dating-question-backend/internal/db"
	"dating-question-backend/internal/seeder"
)

func main() {
	cfg := config.Load()
	db.RunMigrations(cfg)
	db.Connect(cfg)
	seeder.Run(db.DB)
}
