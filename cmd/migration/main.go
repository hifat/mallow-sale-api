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

	log.Println("Starting database migration...")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Create collections and indexes
	if err := CreateUsageUnitCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create usage_units collection: %v", err)
	}

	if err := CreateSettingsCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create settings collection: %v", err)
	}

	if err := CreateRecipeTypeCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create recipe_types collection: %v", err)
	}

	if err := CreateSupplierCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create suppliers collection: %v", err)
	}

	if err := CreateUserCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create users collection: %v", err)
	}

	if err := CreateInventoryCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create inventories collection: %v", err)
	}

	if err := CreateRecipeCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create recipes collection: %v", err)
	}

	if err := CreateStockCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create stocks collection: %v", err)
	}

	if err := CreateShoppingCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create shoppings collection: %v", err)
	}

	if err := CreatePromotionCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create promotions collection: %v", err)
	}

	if err := CreateShoppingInventoryCollection(ctx, db); err != nil {
		log.Fatalf("Failed to create shopping_inventories collection: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
