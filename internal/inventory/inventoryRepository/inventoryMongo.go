package inventoryRepository

import (
	"context"
	"time"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/pkg/database"
	core "github.com/hifat/goroger-core"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *inventoryMongo) Create(ctx context.Context, req inventory.InventoryReq) (string, error) {
	newInventory := inventory.Inventory{}
	if err := copier.Copy(&newInventory, req); err != nil {
		return "", err
	}

	newInventory.SetDateTime()

	_, err := r.db.Collection(newInventory.DocName()).
		InsertOne(ctx, newInventory)

	return newInventory.ID, err
}

func (r *inventoryMongo) Find(ctx context.Context) ([]inventory.Inventory, error) {
	_inventory := inventory.Inventory{}
	cur, err := r.db.Collection(_inventory.DocName()).
		Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	inventories := []inventory.Inventory{}

	return inventories, cur.All(ctx, &inventories)
}

func (r *inventoryMongo) FindByID(ctx context.Context, id string) (*inventory.Inventory, error) {
	_inventory := inventory.Inventory{}
	err := r.db.Collection(_inventory.DocName()).
		FindOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}).Decode(&_inventory)

	return &_inventory, err
}

func (r *inventoryMongo) Update(ctx context.Context, id string, req inventory.InventoryReq) error {
	editInventory := inventory.Inventory{}
	_, err := r.db.Collection(editInventory.DocName()).
		UpdateByID(ctx, database.MustStrToObjectID(id), bson.M{
			"$set": bson.M{
				"name":              req.Name,
				"purchase_price":    req.PurchasePrice,
				"purchase_quantity": req.PurchaseQuantity,
				"purchase_unit": bson.M{
					"code": req.PurchaseUnit.Code,
					"name": req.PurchaseUnit.Name,
				},
				"yield_percentage": req.YieldPercentage,
				"remark":           req.Remark,
				"updated_at":       time.Now(),
			},
		})

	return err
}

func (r *inventoryMongo) Delete(ctx context.Context, id string) error {
	_inventory := inventory.Inventory{}
	_, err := r.db.Collection(_inventory.DocName()).
		DeleteOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		})

	return err
}
