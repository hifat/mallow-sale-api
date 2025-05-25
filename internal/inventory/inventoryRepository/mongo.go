package inventoryRepository

import (
	"context"
	"errors"
	"time"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/throw"
	"github.com/hifat/mallow-sale-api/pkg/utils/repoUtils"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	_, err := r.db.Collection(newInventory.Doc()).
		InsertOne(ctx, newInventory)

	return newInventory.ID, repoUtils.MongoErr(err)
}

func (r *inventoryMongo) Find(ctx context.Context) ([]inventory.Inventory, error) {
	_inventory := inventory.Inventory{}
	cur, err := r.db.Collection(_inventory.Doc()).
		Find(ctx, bson.M{})
	if err != nil {
		return nil, repoUtils.MongoErr(err)
	}
	defer cur.Close(ctx)

	inventories := []inventory.Inventory{}

	return inventories, repoUtils.MongoErr(cur.All(ctx, &inventories))
}

func (r *inventoryMongo) FindByID(ctx context.Context, id string) (*inventory.Inventory, error) {
	_inventory := inventory.Inventory{}
	err := r.db.Collection(_inventory.Doc()).
		FindOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		}).Decode(&_inventory)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, throw.ErrRecordNotFound
	}

	return &_inventory, repoUtils.MongoErr(err)
}

func (r *inventoryMongo) FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.Inventory, error) {
	objectIDs := make([]primitive.ObjectID, 0, len(filter.IDs))
	for _, id := range filter.IDs {
		objectIDs = append(objectIDs, database.MustStrToObjectID(id))
	}

	_inventory := inventory.Inventory{}
	cur, err := r.db.Collection(_inventory.Doc()).
		Find(ctx, bson.M{
			"_id": bson.M{
				"$in": objectIDs,
			},
			"code": bson.M{
				"$in": filter.Codes,
			},
		})
	if err != nil {
		return nil, repoUtils.MongoErr(err)
	}
	defer cur.Close(ctx)

	inventories := []inventory.Inventory{}

	return inventories, repoUtils.MongoErr(cur.All(ctx, &inventories))
}

func (r *inventoryMongo) Update(ctx context.Context, id string, req inventory.InventoryReq) error {
	editInventory := inventory.Inventory{}
	_, err := r.db.Collection(editInventory.Doc()).
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

	return repoUtils.MongoErr(err)
}

func (r *inventoryMongo) Delete(ctx context.Context, id string) error {
	_inventory := inventory.Inventory{}
	_, err := r.db.Collection(_inventory.Doc()).
		DeleteOne(ctx, bson.M{
			"_id": database.MustStrToObjectID(id),
		})

	return repoUtils.MongoErr(err)
}
