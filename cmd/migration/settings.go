package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSettingsCollection(ctx context.Context, db *mongo.Database) error {
	collectionName := "settings"

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

	return nil
}
