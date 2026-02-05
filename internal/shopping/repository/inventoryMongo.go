package shoppingRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
)

type inventoryMongoRepository struct {
	db *mongo.Database
}

func NewInventoryMongo(db *mongo.Database) shoppingModule.IInventoryRepository {
	return &inventoryMongoRepository{db: db}
}

func (r *inventoryMongoRepository) Create(ctx context.Context, req *shoppingModule.RequestShoppingInventory) error {
	newInvShopping := shoppingModule.InventoryEntity{
		InventoryID:   req.InventoryID,
		InventoryName: req.InventoryName,
		SupplierID:    req.SupplierID,
		SupplierName:  req.SupplierName,
	}

	_, err := r.db.Collection("shopping_inventories").InsertOne(ctx, newInvShopping)
	if err != nil {
		return err
	}

	return nil
}

func (r *inventoryMongoRepository) Find(ctx context.Context) ([]shoppingModule.ResShoppingInventory, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$supplier_id"},
			{Key: "supplierName", Value: bson.D{{Key: "$first", Value: "$supplier_name"}}},
			{Key: "inventories", Value: bson.D{{Key: "$push", Value: bson.D{
				{Key: "id", Value: "$_id"},
				{Key: "inventoryID", Value: "$inventory_id"},
				{Key: "inventoryName", Value: "$inventory_name"},
			}}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "supplierID", Value: "$supplier_id"},
			{Key: "supplierName", Value: 1},
			{Key: "inventories", Value: 1},
		}}},
	}

	cur, err := r.db.Collection("shopping_inventories").Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var responses []shoppingModule.ResShoppingInventory
	if err := cur.All(ctx, &responses); err != nil {
		return nil, err
	}

	return responses, nil
}

func (r *inventoryMongoRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.Collection("shopping_inventories").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
