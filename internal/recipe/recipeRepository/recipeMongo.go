package recipeRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/internal/usageUnit"
	"github.com/hifat/cost-calculator-api/pkg/database"
	core "github.com/hifat/goroger-core"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeMongo struct {
	db     *mongo.Database
	helper core.Helper
}

func NewMongo(db *mongo.Database, helper core.Helper) IRecipeRepository {
	return &recipeMongo{
		db,
		helper,
	}
}

func (r *recipeMongo) Create(ctx context.Context, req recipe.RecipeReq) (id string, err error) {
	newRecipe := new(recipe.Recipe)
	if err = r.helper.Copy(newRecipe, req); err != nil {
		return "", err
	}

	newRecipe.SetDateTime()

	for i := range newRecipe.Inventories {
		newRecipe.Inventories[i].SetDateTime()
		newRecipe.Inventories[i].SetID()
		newRecipe.Inventories[i].UsageUnit = &usageUnit.UsageUnitEmbed{
			Code: req.Inventories[i].UsageUnitCode,
		}
	}

	result, err := r.db.Collection(newRecipe.DocName()).
		InsertOne(ctx, newRecipe)
	if err != nil {
		return "", err
	}

	latestID, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return latestID.Hex(), nil
	}

	return "none-id", nil
}

func (r *recipeMongo) Find(ctx context.Context) ([]recipe.RecipeRes, error) {
	_recipe := recipe.Recipe{}
	cur, err := r.db.Collection(_recipe.DocName()).
		Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	recipes := []recipe.Recipe{}
	if err := cur.All(ctx, &recipes); err != nil {
		return nil, err
	}

	res := make([]recipe.RecipeRes, 0, len(recipes))
	if err := r.helper.Copy(&res, recipes); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *recipeMongo) FindByID(ctx context.Context, id string) (*recipe.RecipeRes, error) {
	_recipe := recipe.Recipe{}

	pipeline := mongo.Pipeline{
		{
			{
				Key: "$match", Value: bson.M{
					"_id": database.MustStrToObjectID(id),
				},
			},
		},
		{
			{
				Key: "$lookup", Value: bson.M{
					"from":         "inventories",
					"localField":   "inventories.inventoryID",
					"foreignField": "_id",
					"as":           "inventoryDetails",
				},
			},
		},
		{
			{
				Key: "$lookup", Value: bson.M{
					"from":         "usage_units",
					"localField":   "inventories.usageUnitCode",
					"foreignField": "code",
					"as":           "usageUnitDetails",
				},
			},
		},
	}

	cursor, err := r.db.Collection(_recipe.DocName()).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []recipe.Recipe
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	res := new(recipe.RecipeRes)
	if err := r.helper.Copy(&res, results[0]); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *recipeMongo) Update(ctx context.Context, id string, req recipe.RecipeReq) error {
	editRecipe := recipe.Recipe{}
	editRecipe.SetDateTime()

	if err := copier.Copy(&editRecipe, req); err != nil {
		return err
	}

	for i := range editRecipe.Inventories {
		editRecipe.Inventories[i].SetUpdatedAt()

		if editRecipe.Inventories[i].ID == "" {
			editRecipe.Inventories[i].SetID()
		}
	}

	_, err := r.db.Collection(editRecipe.DocName()).
		UpdateOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}, bson.M{
			"$set": editRecipe,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeMongo) Delete(ctx context.Context, id string) error {
	_recipe := recipe.Recipe{}
	_, err := r.db.Collection(_recipe.DocName()).
		DeleteOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		})

	return err
}
