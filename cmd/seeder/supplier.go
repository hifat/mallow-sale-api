package main

import (
	"context"
	"log"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedSuppliers(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("suppliers")

	// Define initial suppliers
	usageUnits := []supplierModule.Entity{
		{
			Name:   "Internal",
			ImgUrl: "",
		},
		{
			Name:   "Makro",
			ImgUrl: "",
		},
		{
			Name:   "กาดเมืองใหม่",
			ImgUrl: "",
		},
		{
			Name:   "กาดแม่โจ้",
			ImgUrl: "",
		},
		{
			Name:   "Lotus",
			ImgUrl: "",
		},
	}

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Println("Suppliers already seeded, skipping...")
		return nil
	}

	// Convert to interface slice for bulk insert
	var documents []interface{}
	for _, unit := range usageUnits {
		documents = append(documents, unit)
	}

	// Insert all suppliers
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	log.Printf("Successfully seeded %d suppliers", len(result.InsertedIDs))
	return nil
}
