package inventoryRepository

import (
	"context"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/pkg/database"
	"github.com/hifat/cost-calculator-api/pkg/utils"
	core "github.com/hifat/goroger-core"
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

func (r *inventoryMongo) Create(ctx context.Context, req inventory.Inventory) (string, error) {
	req.SetDateTime()

	_, err := r.db.Collection(req.DocName()).
		InsertOne(ctx, req)

	return req.ID, err
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

func (r *inventoryMongo) Update(ctx context.Context, id string, req inventory.Inventory) error {
	_inventory := inventory.Inventory{}
	req.UpdatedAt = utils.TimeNow()
	_, err := r.db.Collection(_inventory.DocName()).
		UpdateOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}, bson.M{
			"$set": req,
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
