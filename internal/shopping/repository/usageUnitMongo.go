package shoppingRepository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
)

type usageUnitMongoRepository struct {
	db *mongo.Database
}

func NewUsageUnitMongo(db *mongo.Database) shoppingModule.IUsageUnitRepository {
	return &usageUnitMongoRepository{db: db}
}

func (r *usageUnitMongoRepository) Create(ctx context.Context, req *shoppingModule.RequestUsageUnit) error {
	newUsageUnit := shoppingModule.UsageUnitEntity{
		Code: req.Code,
		Name: req.Name,
	}

	_, err := r.db.Collection("shopping_usage_units").
		InsertOne(ctx, newUsageUnit)
	if err != nil {
		return err
	}

	return nil
}

func (r *usageUnitMongoRepository) Find(ctx context.Context) ([]shoppingModule.ResUsageUnit, error) {
	cursor, err := r.db.Collection("shopping_usage_units").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var usageUnits []shoppingModule.ResUsageUnit
	if err := cursor.All(ctx, &usageUnits); err != nil {
		return nil, err
	}

	return usageUnits, nil
}

func (r *usageUnitMongoRepository) FindByID(ctx context.Context, id string) (*shoppingModule.ResUsageUnit, error) {
	cursor := r.db.Collection("shopping_usage_units").FindOne(ctx, bson.M{"_id": database.MustObjectIDFromHex(id)})
	if cursor.Err() != nil {
		if errors.Is(cursor.Err(), mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, cursor.Err()
	}

	var usageUnit shoppingModule.ResUsageUnit
	if err := cursor.Decode(&usageUnit); err != nil {
		return nil, err
	}

	return &usageUnit, nil
}

func (r *usageUnitMongoRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.Collection("shopping_usage_units").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *usageUnitMongoRepository) UpdateByID(ctx context.Context, id string, req *shoppingModule.RequestUsageUnit) error {
	_, err := r.db.Collection("shopping_usage_units").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"code": req.Code, "name": req.Name}})
	if err != nil {
		return err
	}

	return nil
}
