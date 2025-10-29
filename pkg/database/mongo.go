package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(cfg config.DB) (*mongo.Database, func(), error) {
	// Build connection string with proper authentication
	uri := fmt.Sprintf(
		"%s://%s:%s@%s%s/%s",
		cfg.Protocol,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		func() string {
			if cfg.Port != "" {
				return ":" + cfg.Port
			}

			return ""
		}(),
		cfg.DBName,
	)

	if cfg.Query != "" {
		uri += cfg.Query
	}

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

func MustObjectIDFromHex(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Warn("Invalid ObjectID", "id", id, "error", err)
		return primitive.NilObjectID
	}

	return objectID
}
