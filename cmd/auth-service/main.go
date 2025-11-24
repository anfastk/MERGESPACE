package main

import (
	"log"

	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/config"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/database"
	"github.com/anfastk/MERGESPACE/internal/auth-service/infrastructure/di"
)

func main() {
	
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	appServer, err := di.InitializeApp(db, cfg)
	if err != nil {
		log.Fatal(err)
	}
	appServer.Run()
}
