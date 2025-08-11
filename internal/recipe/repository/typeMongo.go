package recipeRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type mongoTypeRepository struct {
	db *mongo.Database
}

func NewTypeMongo(db *mongo.Database) TypeRepository {
	return &mongoTypeRepository{db: db}
}

func (r *mongoTypeRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]recipeModule.RecipeTypeResponse, error) {
	cursor, err := r.db.Collection("recipe_types").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var types []recipeModule.RecipeTypeResponse
	for cursor.Next(ctx) {
		var typeResponse recipeModule.RecipeTypeResponse
		if err := cursor.Decode(&typeResponse); err != nil {
			return nil, err
		}
		types = append(types, typeResponse)
	}

	return types, nil
}

func (r *mongoTypeRepository) FindByCode(ctx context.Context, code string) (*recipeModule.RecipeTypeResponse, error) {
	result := r.db.Collection("recipe_types").FindOne(ctx, bson.M{"code": code})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var typeResponse recipeModule.RecipeTypeResponse
	if err := result.Decode(&typeResponse); err != nil {
		return nil, err
	}

	return &typeResponse, nil
}

func (r *mongoTypeRepository) FindInCodes(ctx context.Context, codes []string) ([]recipeModule.RecipeTypeResponse, error) {
	cursor, err := r.db.Collection("recipe_types").Find(ctx, bson.M{"code": bson.M{"$in": codes}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var types []recipeModule.RecipeTypeResponse
	for cursor.Next(ctx) {
		var typeResponse recipeModule.RecipeTypeResponse
		if err := cursor.Decode(&typeResponse); err != nil {
			return nil, err
		}
		types = append(types, typeResponse)
	}

	return types, nil
}
