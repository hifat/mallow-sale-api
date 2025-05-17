package usageUnitRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type usageUnitMongo struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IUsageUnitRepository {
	return &usageUnitMongo{db}
}

func (r *usageUnitMongo) FindInCodes(ctx context.Context, codes []string) ([]usageUnit.UsageUnit, error) {
	filter := bson.M{
		"code": bson.M{
			"$in": codes,
		},
	}

	_usageUnit := usageUnit.UsageUnit{}
	cur, err := r.db.Collection(_usageUnit.DocName()).
		Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	units := []usageUnit.UsageUnit{}

	return units, cur.All(ctx, &units)
}
