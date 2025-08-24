package main

import (
	"context"
	"log"
	"time"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./env/.env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, cleanup, err := database.NewMongo(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer cleanup()

	log.Println("Starting database seeding...")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Seed usage units
	if err := SeedUsageUnits(ctx, db); err != nil {
		log.Fatalf("Failed to seed usage units: %v", err)
	}

	// Seed settings
	if err := SeedSettings(ctx, db); err != nil {
		log.Fatalf("Failed to seed settings: %v", err)
	}

	// Seed recipe types
	if err := SeedRecipeTypes(ctx, db); err != nil {
		log.Fatalf("Failed to seed recipe types: %v", err)
	}

	// Seed suppliers
	if err := SeedSuppliers(ctx, db); err != nil {
		log.Fatalf("Failed to seed suppliers: %v", err)
	}

	log.Println("Database seeding completed successfully!")
}
