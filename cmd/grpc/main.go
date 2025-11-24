package main

import (
	"flag"
	"log"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/rpc"
)

func main() {
	envPath := flag.String("envPath", "", "env path")
	flag.Parse()

	cfg, err := config.LoadConfig(*envPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, cleanup, err := database.NewMongo(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer cleanup()

	rpc.RegisterGRPC(cfg, db)
}
