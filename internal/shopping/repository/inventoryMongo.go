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

func (r *inventoryMongoRepository) Find(ctx context.Context) ([]shoppingModule.InventoryResponse, error) {
	cur, err := r.db.Collection("shopping_inventories").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var invShoppings []shoppingModule.InventoryEntity
	if err := cur.All(ctx, &invShoppings); err != nil {
		return nil, err
	}

	var responses []shoppingModule.InventoryResponse
	for _, inv := range invShoppings {
		responses = append(responses, shoppingModule.InventoryResponse{
			ID:            inv.ID.Hex(),
			InventoryID:   inv.InventoryID,
			InventoryName: inv.InventoryName,
			SupplierID:    inv.SupplierID,
			SupplierName:  inv.SupplierName,
		})
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
