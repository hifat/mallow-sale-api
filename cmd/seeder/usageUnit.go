package main

import (
	"context"
	"log"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedUsageUnits(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("usage_units")

	// Define initial usage units
	usageUnits := []usageUnitModule.Entity{
		{Code: "kg", Name: "กิโลกรัม"},
		{Code: "g", Name: "กรัม"},
		{Code: "l", Name: "ลิตร"},
		{Code: "ml", Name: "มิลลิลิตร"},
		{Code: "pcs", Name: "ชิ้น"},
		{Code: "box", Name: "กล่อง"},
		{Code: "pack", Name: "แพ็ค"},
		{Code: "bottle", Name: "ขวด"},
		{Code: "can", Name: "กระป๋อง"},
		{Code: "bag", Name: "ถุง"},
		{Code: "jar", Name: "โหล"},
		{Code: "cup", Name: "ถ้วย"},
		{Code: "tbsp", Name: "ช้อนโต๊ะ"},
		{Code: "tsp", Name: "ช้อนชา"},
		{Code: "oz", Name: "ออนซ์"},
		{Code: "lb", Name: "ปอนด์"},
		{Code: "gal", Name: "แกลลอน"},
		{Code: "qt", Name: "ควอร์ต"},
		{Code: "pt", Name: "ไพน์ต"},
		{Code: "fl_oz", Name: "ออนซ์ของเหลว"},
	}

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Println("Usage units already seeded, skipping...")
		return nil
	}

	// Convert to interface slice for bulk insert
	var documents []interface{}
	for _, unit := range usageUnits {
		documents = append(documents, unit)
	}

	// Insert all usage units
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	log.Printf("Successfully seeded %d usage units", len(result.InsertedIDs))
	return nil
}
