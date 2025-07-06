package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Seeder struct {
	collection *mongo.Collection
}

func NewSeeder(collection *mongo.Collection) *Seeder {
	return &Seeder{
		collection: collection,
	}
}

// Seed runs all seeding operations for usage units
func (s *Seeder) Seed(ctx context.Context) error {
	log.Println("Starting recipe seeding...")

	log.Println("Recipe seeding completed successfully")
	return nil
}

// SeedIfEmpty seeds the database only if it's empty
func (s *Seeder) SeedIfEmpty(ctx context.Context) error {
	count, err := s.collection.CountDocuments(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	if count == 0 {
		return s.Seed(ctx)
	}

	log.Println("Recipes collection is not empty, skipping seeding")
	return nil
}
