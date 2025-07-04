package database

import (
	"context"
	"fmt"
	"time"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(cfg config.DB) (*mongo.Database, func(), error) {
	// Build connection string with proper authentication
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	if cfg.Schema != "" {
		uri += fmt.Sprintf("?authSource=%s", cfg.Schema)
	} else {
		uri += "?authSource=admin"
	}

	if cfg.SSLMode != "" {
		uri += fmt.Sprintf("&ssl=%s", cfg.SSLMode)
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
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Successfully connected to MongoDB")

	return client.Database(cfg.DBName), func() {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Printf("Error disconnecting from MongoDB: %v\n", err)
		}
	}, nil
}
