package inventoryRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	core "github.com/hifat/goroger-core"
	"go.mongodb.org/mongo-driver/mongo"
)

type inventoryMongo struct {
	db     *mongo.Database
	helper core.Helper
}

func NewMongo(db *mongo.Database, helper core.Helper) IInventoryRepository {
	return &inventoryMongo{
		db,
		helper,
	}
}

func (r *inventoryMongo) Create(ctx context.Context, req inventory.Inventory) (string, error) {
	newInventory := inventory.InventoryEntity{}
	if err := r.helper.Copy(&newInventory, req); err != nil {
		return "", err
	}
	newInventory.SetDateTime()

	_, err := r.db.Collection(newInventory.DocName()).
		InsertOne(ctx, newInventory)
	if err != nil {
		return "", err
	}

	return req.ID, nil
}

func (r *inventoryMongo) Find(ctx context.Context) ([]inventory.Inventory, error) {
	_inventory := inventory.InventoryEntity{}
	cur, err := r.db.Collection(_inventory.DocName()).
		Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	inventories := []inventory.InventoryEntity{}
	if err := cur.All(ctx, &inventories); err != nil {
		return nil, err
	}

	res := []inventory.Inventory{}
	if err := r.helper.Copy(&res, inventories); err != nil {
		return nil, err
	}

	return res, nil
}
