package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SeedSettings(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("settings")

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Println("Settings already seeded, skipping...")
		return nil
	}

	update := bson.M{"$setOnInsert": bson.M{"cost_percentage": 30.00}} // Default value
	isUpsert := true
	_, err = collection.UpdateOne(ctx, bson.M{}, update, &options.UpdateOptions{Upsert: &isUpsert})
	if err != nil {
		return err
	}
	log.Println("Settings seeded (if not present)")
	return nil
}
