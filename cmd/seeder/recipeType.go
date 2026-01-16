package main

import (
	"context"
	"log"

	recipeTypeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedRecipeTypes(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection("recipe_types")

	// Define initial recipe types
	recipeTypes := []recipeTypeModule.RecipeTypeEntity{
		{Code: "FOOD", Name: "Food"},
		{Code: "DESSERT", Name: "Dessert"},
		{Code: "DRINK", Name: "Drink"},
		{Code: "SNACK", Name: "Snack"},
		{Code: "INGREDIENT", Name: "Ingredient"},
	}

	// Check if data already exists
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// If data already exists, skip seeding
	if count > 0 {
		log.Println("Recipe types already seeded, skipping...")
		return nil
	}

	// Convert to interface slice for bulk insert
	var documents []interface{}
	for _, recipeType := range recipeTypes {
		documents = append(documents, recipeType)
	}

	// Insert all recipe types
	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return err
	}

	log.Printf("Successfully seeded %v recipe types", len(result.InsertedIDs))
	return nil
}
