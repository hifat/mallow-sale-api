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

func NewMongo(db *mongo.Database) settingModule.Repository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Update(costPercentage float32) error {
	filter := bson.M{"_id": "global"}
	update := bson.M{"$set": bson.M{"cost_percentage": costPercentage}}
	_, err := r.db.Collection("settings").UpdateOne(context.Background(), filter, update, &options.UpdateOptions{Upsert: new(bool)})
	return err
}

func (r *mongoRepository) Get() (*settingModule.Entity, error) {
	filter := bson.M{"_id": "global"}
	var settings settingModule.Entity
	err := r.db.Collection("settings").FindOne(context.Background(), filter).Decode(&settings)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &settingModule.Entity{CostPercentage: 0}, nil
		}
		return nil, err
	}
	return &settings, nil
}
