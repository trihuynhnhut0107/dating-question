package main

import (
	"dating-question-backend/internal/config"
	"dating-question-backend/internal/db"
	"dating-question-backend/internal/handler"
)

func main() {
	cfg := config.Load()
	db.RunMigrations(cfg)
	db.Connect(cfg)

	router := handler.SetupRouter()
	router.Run(cfg.ServerPort)
}
