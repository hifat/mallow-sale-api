package main

import (
	"context"
	"log"
	"time"

	userModule "github.com/hifat/mallow-sale-api/internal/user"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedUser(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("users")

	// Define initial user
	newUser := userModule.Entity{
		Base: utilsModule.Base{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "admin",
		Username: "admin",
		Password: "$2a$10$lNVnjXjniMR6BTZzneN6TuvB1n/DKl.DcoS.aa4WBxO7XLEzy.QDq", // Default: 1234
	}

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Println("User already seeded, skipping...")
		return nil
	}

	// Insert all user
	result, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	log.Printf("Successfully seeded %d user", result.InsertedID)
	return nil
}
