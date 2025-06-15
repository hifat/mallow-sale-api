package recipeRepository

import (
	"context"
	"time"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/internal/recipe"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/utils/repoUtils"
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

	for i := range newRecipe.Ingredients {
		newRecipe.Ingredients[i].ID = primitive.NewObjectID().Hex()
	}

	result, err := r.db.Collection(newRecipe.Doc()).
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
	cur, err := r.db.Collection(_recipe.Doc()).
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
	err := r.db.Collection(_recipe.Doc()).
		FindOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}).Decode(&_recipe)
	if err != nil {
		return nil, repoUtils.MongoErr(err)
	}

	res := new(recipe.RecipeRes)
	if err := r.helper.Copy(&res, _recipe); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *recipeMongo) Update(ctx context.Context, id string, req recipe.UpdateRecipeReq) error {
	editRecipe := recipe.Recipe{}
	editRecipe.SetDateTime()

	if err := copier.Copy(&editRecipe, req); err != nil {
		return err
	}

	setIngredients := make([]bson.M, 0, len(req.Ingredients))
	for _, _inventory := range editRecipe.Ingredients {
		code, name := "", ""
		if _inventory.UsageUnit != nil {
			code = _inventory.UsageUnit.Code
			name = _inventory.UsageUnit.Name
		}

		mapInventory := bson.M{
			"_id":            primitive.NewObjectID().Hex(),
			"usage_quantity": _inventory.UsageQuantity,
			"remark":         _inventory.Remark,
			"usage_unit": bson.M{
				"code": code,
				"name": name,
			},
			"inventory_id": _inventory.InventoryID,
		}

		setIngredients = append(setIngredients, mapInventory)
	}

	setRecipe := bson.M{
		"name":        req.Name,
		"updated_at":  time.Now(),
		"ingredients": setIngredients,
	}

	_, err := r.db.Collection(editRecipe.Doc()).
		UpdateOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}, bson.M{
			"$set": setRecipe,
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeMongo) Delete(ctx context.Context, id string) error {
	_recipe := recipe.Recipe{}
	_, err := r.db.Collection(_recipe.Doc()).
		DeleteOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		})

	return err
}
