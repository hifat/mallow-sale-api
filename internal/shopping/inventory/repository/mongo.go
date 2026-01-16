package shoppingInventoryRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/shopping/inventory"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) inventoryModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *inventoryModule.Request) error {
	newInvShoping := inventoryModule.Entity{
		InventoryID:   req.InventoryID,
		InventoryName: req.InventoryName,
		SupplierID:    req.SupplierID,
		SupplierName:  req.SupplierName,
	}

	_, err := r.db.Collection("inventory").InsertOne(ctx, newInvShoping)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Find(ctx context.Context) ([]inventoryModule.Response, error) {
	cursor, err := r.db.Collection("inventory").Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	var invShopings []inventoryModule.Response
	if err := cursor.All(ctx, &invShopings); err != nil {
		return nil, err
	}

	return invShopings, nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.Collection("inventory").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
