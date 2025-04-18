package recipeRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/recipe"
	"github.com/hifat/cost-calculator-api/pkg/database"
	core "github.com/hifat/goroger-core"
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

	result, err := r.db.Collection(newRecipe.DocName()).
		InsertOne(ctx, req)
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
	err := r.db.Collection(_recipe.DocName()).
		FindOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}).Decode(&_recipe)
	if err != nil {
		return nil, err
	}

	res := new(recipe.RecipeRes)
	if err := r.helper.Copy(&res, _recipe); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *recipeMongo) Update(ctx context.Context, id string, req recipe.RecipeReq) error {
	_recipe := recipe.Recipe{}
	_recipe.SetDateTime()
	_, err := r.db.Collection(_recipe.DocName()).
		UpdateOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}, bson.M{
			"$set": req,
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
