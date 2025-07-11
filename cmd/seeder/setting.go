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
	update := bson.M{"$setOnInsert": bson.M{"cost_percentage": 30.0}} // Default value
	_, err := collection.UpdateOne(ctx, bson.M{}, update, &options.UpdateOptions{Upsert: new(bool)})
	if err != nil {
		return err
	}
	log.Println("Settings seeded (if not present)")
	return nil
}
