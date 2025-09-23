package settingRepository

import (
	"context"

	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Find(ctx context.Context) (*settingModule.Response, error) {
	var settings settingModule.Entity
	err := r.db.Collection("settings").FindOne(ctx, bson.M{}).Decode(&settings)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &settingModule.Response{CostPercentage: 0}, nil
		}

		return nil, err
	}

	return &settingModule.Response{
		CostPercentage: settings.CostPercentage,
	}, nil
}

func (r *mongoRepository) Update(ctx context.Context, costPercentage float32) error {
	update := bson.M{"$set": bson.M{"cost_percentage": costPercentage}}
	_, err := r.db.Collection("settings").UpdateOne(ctx, bson.M{}, update, &options.UpdateOptions{Upsert: new(bool)})
	return err
}
