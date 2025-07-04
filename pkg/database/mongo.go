package database

import (
	"context"
	"fmt"
	"time"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(cfg config.DB) (*mongo.Database, error) {
	// Build connection string with proper authentication
	var uri string

	if cfg.Username != "" && cfg.Password != "" {
		// With authentication
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)

		// Add authSource parameter
		if cfg.Schema != "" {
			uri += fmt.Sprintf("?authSource=%s", cfg.Schema)
		} else {
			uri += "?authSource=admin"
		}

		// Add SSL mode if specified
		if cfg.SSLMode != "" {
			if cfg.Schema != "" {
				uri += fmt.Sprintf("&ssl=%s", cfg.SSLMode)
			} else {
				uri += fmt.Sprintf("&ssl=%s", cfg.SSLMode)
			}
		}
	} else {
		// Without authentication (for local development)
		uri = fmt.Sprintf(
			"mongodb://%s:%s/%s",
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)

		// Add SSL mode if specified
		if cfg.SSLMode != "" {
			uri += fmt.Sprintf("?ssl=%s", cfg.SSLMode)
		}
	}

	fmt.Printf("Connecting to MongoDB: %s\n", uri)

	clientOpts := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(10 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client.Database(cfg.DBName), nil
}
