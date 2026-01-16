package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRecipeTypeCollection(ctx context.Context, db *mongo.Database) error {
	collectionName := "recipe_types"

	// Create collection if it doesn't exist
	err := db.CreateCollection(ctx, collectionName)
	if err != nil {
		// Check if collection already exists
		if mongo.IsDuplicateKeyError(err) || err.Error() == "collection already exists" {
			log.Printf("Collection '%s' already exists, skipping creation...", collectionName)
		} else {
			return err
		}
	} else {
		log.Printf("Created collection: %s", collectionName)
	}

	collection := db.Collection(collectionName)

	// Create indexes
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
	}

	_, err = collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("Warning: Failed to create indexes for %s: %v", collectionName, err)
	} else {
		log.Printf("Created indexes for: %s", collectionName)
	}

	return nil
}
