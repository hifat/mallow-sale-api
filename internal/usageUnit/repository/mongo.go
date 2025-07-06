package usageUnitRepository

import (
	"context"
	"errors"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) FindByCode(ctx context.Context, code string) (*usageUnitModule.Prototype, error) {
	filter := bson.M{"code": code}
	var usageUnit usageUnitModule.Entity
	err := r.db.Collection("usage_units").
		FindOne(ctx, filter).
		Decode(&usageUnit)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	return &usageUnitModule.Prototype{
		Code: usageUnit.Code,
		Name: usageUnit.Name,
	}, nil
}

func (r *mongoRepository) FindInCodes(ctx context.Context, codes []string) ([]usageUnitModule.Prototype, error) {
	filter := bson.M{"code": bson.M{"$in": codes}}
	usageUnits := make([]usageUnitModule.Prototype, 0)
	cursor, err := r.db.Collection("usage_units").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var usageUnit usageUnitModule.Entity
		if err := cursor.Decode(&usageUnit); err != nil {
			return nil, err
		}

		usageUnits = append(usageUnits, usageUnitModule.Prototype(usageUnit))
	}

	return usageUnits, nil
}
