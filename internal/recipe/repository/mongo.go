package recipeRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *recipeModule.Request) error {
	ingredients := make([]recipeModule.IngredientEntity, 0, len(req.Ingredients))
	for _, ingredient := range req.Ingredients {
		ingredients = append(ingredients, recipeModule.IngredientEntity{
			InventoryID: database.MustObjectIDFromHex(ingredient.InventoryID),
			Quantity:    ingredient.Quantity,
			Unit: usageUnitModule.Entity{
				Code: ingredient.Unit.Code,
				Name: ingredient.Unit.Name,
			},
		})
	}

	newRecipe := &recipeModule.Entity{
		Name:        req.Name,
		Ingredients: ingredients,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := r.db.Collection("recipes").InsertOne(ctx, newRecipe)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]recipeModule.Response, error) {
	filter := bson.M{}
	if query.Search != "" {
		filter["name"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	findOptions := options.Find()

	if query.Sort != "" && query.Order != "" {
		order := 1
		if strings.ToLower(query.Order) == "desc" {
			order = -1
		}
		findOptions.SetSort(bson.M{query.Sort: order})
	}

	if query.Page > 0 && query.Limit > 0 {
		findOptions.SetSkip(int64((query.Page - 1) * query.Limit))
		findOptions.SetLimit(int64(query.Limit))
	}

	if query.Fields != "" {
		fields := strings.Split(query.Fields, ",")
		projection := bson.M{}
		for _, field := range fields {
			projection[field] = 1
		}

		findOptions.SetProjection(projection)
	}

	recipes := make([]recipeModule.Entity, 0)
	cursor, err := r.db.Collection("recipes").Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var recipe recipeModule.Entity
		if err := cursor.Decode(&recipe); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	res := make([]recipeModule.Response, 0, len(recipes))
	for _, recipe := range recipes {
		ingredients := make([]recipeModule.IngredientPrototype, 0, len(recipe.Ingredients))
		for _, ingredient := range recipe.Ingredients {
			ingredients = append(ingredients, recipeModule.IngredientPrototype{
				InventoryID: ingredient.InventoryID.Hex(),
				Quantity:    ingredient.Quantity,
				Unit: usageUnitModule.Prototype{
					Code: ingredient.Unit.Code,
					Name: ingredient.Unit.Name,
				},
			})
		}

		res = append(res, recipeModule.Response{
			Prototype: recipeModule.Prototype{
				ID:          recipe.ID.Hex(),
				Name:        recipe.Name,
				Ingredients: ingredients,
				CreatedAt:   &recipe.CreatedAt,
				UpdatedAt:   &recipe.UpdatedAt,
			},
		})
	}

	return res, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*recipeModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var recipe recipeModule.Entity
	err := r.db.Collection("recipes").
		FindOne(ctx, filter).
		Decode(&recipe)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	ingredients := make([]recipeModule.IngredientPrototype, 0, len(recipe.Ingredients))
	for _, ingredient := range recipe.Ingredients {
		ingredients = append(ingredients, recipeModule.IngredientPrototype{
			InventoryID: ingredient.InventoryID.Hex(),
			Quantity:    ingredient.Quantity,
			Unit: usageUnitModule.Prototype{
				Code: ingredient.Unit.Code,
				Name: ingredient.Unit.Name,
			},
		})
	}

	return &recipeModule.Response{
		Prototype: recipeModule.Prototype{
			ID:          recipe.ID.Hex(),
			Name:        recipe.Name,
			Ingredients: ingredients,
			CreatedAt:   &recipe.CreatedAt,
			UpdatedAt:   &recipe.UpdatedAt,
		},
	}, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *recipeModule.Request) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	editedRecipe := &recipeModule.Entity{
		Name: req.Name,
		Base: utilsModule.Base{
			UpdatedAt: time.Now(),
		},
	}

	ingredients := make([]recipeModule.IngredientEntity, len(req.Ingredients))
	for i, ingredient := range req.Ingredients {
		ingredients[i] = recipeModule.IngredientEntity{
			InventoryID: database.MustObjectIDFromHex(ingredient.InventoryID),
			Quantity:    ingredient.Quantity,
			Unit: usageUnitModule.Entity{
				Code: ingredient.Unit.Code,
				Name: ingredient.Unit.Name,
			},
		}
	}
	editedRecipe.Ingredients = ingredients

	_, err := r.db.Collection("recipes").UpdateOne(ctx, filter, bson.M{"$set": editedRecipe})
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("recipes").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Collection("recipes").CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
