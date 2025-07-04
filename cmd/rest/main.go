package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/router"
)

func main() {
	cfg, err := config.LoadConfig("./env/.env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMongo(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := gin.Default()
	router.RegisterAll(r, cfg, db)

	r.Run(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port))
}
