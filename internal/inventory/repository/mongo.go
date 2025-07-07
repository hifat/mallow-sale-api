package inventoryRepository

import (
	"context"
	"errors"
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) Repository {
	return &mongoRepository{
		db: db,
	}
}

func (r *mongoRepository) Create(ctx context.Context, req *inventoryModule.Request) error {
	newInventory := &inventoryModule.Entity{
		Name:             req.Name,
		PurchasePrice:    req.PurchasePrice,
		PurchaseQuantity: req.PurchaseQuantity,
		PurchaseUnit: usageUnitModule.Entity{
			Code: req.PurchaseUnit.Code,
			Name: req.PurchaseUnit.Name,
		},
		YieldPercentage: req.YieldPercentage,
		Remark:          req.Remark,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := r.db.Collection("inventories").
		InsertOne(ctx, newInventory)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*inventoryModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var inventory inventoryModule.Entity
	err := r.db.Collection("inventories").
		FindOne(ctx, filter).
		Decode(&inventory)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	res := &inventoryModule.Response{
		Prototype: inventoryModule.Prototype{
			ID:               inventory.ID.Hex(),
			Name:             inventory.Name,
			PurchasePrice:    inventory.PurchasePrice,
			PurchaseQuantity: inventory.PurchaseQuantity,
			PurchaseUnit: usageUnitModule.Prototype{
				Code: inventory.PurchaseUnit.Code,
				Name: inventory.PurchaseUnit.Name,
			},
			YieldPercentage: inventory.YieldPercentage,
			Remark:          inventory.Remark,
			CreatedAt:       &inventory.CreatedAt,
			UpdatedAt:       &inventory.UpdatedAt,
		},
	}

	return res, nil
}

func (r *mongoRepository) Find(ctx context.Context) ([]inventoryModule.Response, error) {
	var inventories []inventoryModule.Entity
	cur, err := r.db.Collection("inventories").
		Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var inventory inventoryModule.Entity
		if err := cur.Decode(&inventory); err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	res := make([]inventoryModule.Response, len(inventories))
	for i, inventory := range inventories {
		res[i] = inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				ID:               inventory.ID.Hex(),
				Name:             inventory.Name,
				PurchasePrice:    inventory.PurchasePrice,
				PurchaseQuantity: inventory.PurchaseQuantity,
				PurchaseUnit: usageUnitModule.Prototype{
					Code: inventory.PurchaseUnit.Code,
					Name: inventory.PurchaseUnit.Name,
				},
				YieldPercentage: inventory.YieldPercentage,
				Remark:          inventory.Remark,
				CreatedAt:       &inventory.CreatedAt,
				UpdatedAt:       &inventory.UpdatedAt,
			},
		}
	}

	return res, nil
}

func (r *mongoRepository) FindInIDs(ctx context.Context, ids []string) ([]inventoryModule.Response, error) {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIDs[i] = database.MustObjectIDFromHex(id)
	}

	filter := bson.M{"_id": bson.M{"$in": objIDs}}
	var inventories []inventoryModule.Entity
	cur, err := r.db.Collection("inventories").
		Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var inventory inventoryModule.Entity
		if err := cur.Decode(&inventory); err != nil {
			return nil, err
		}
		inventories = append(inventories, inventory)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	res := make([]inventoryModule.Response, 0, len(inventories))
	for _, inventory := range inventories {
		res = append(res, inventoryModule.Response{
			Prototype: inventoryModule.Prototype{
				ID:               inventory.ID.Hex(),
				Name:             inventory.Name,
				PurchasePrice:    inventory.PurchasePrice,
				PurchaseQuantity: inventory.PurchaseQuantity,
				PurchaseUnit: usageUnitModule.Prototype{
					Code: inventory.PurchaseUnit.Code,
					Name: inventory.PurchaseUnit.Name,
				},
				YieldPercentage: inventory.YieldPercentage,
				Remark:          inventory.Remark,
				CreatedAt:       &inventory.CreatedAt,
				UpdatedAt:       &inventory.UpdatedAt,
			},
		})
	}

	return res, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *inventoryModule.Request) error {
	editedInventory := &inventoryModule.Entity{
		Name:             req.Name,
		PurchasePrice:    req.PurchasePrice,
		PurchaseQuantity: req.PurchaseQuantity,
		PurchaseUnit: usageUnitModule.Entity{
			Code: req.PurchaseUnit.Code,
			Name: req.PurchaseUnit.Name,
		},
		YieldPercentage: req.YieldPercentage,
		Remark:          req.Remark,
		Base: utilsModule.Base{
			UpdatedAt: time.Now(),
		},
	}

	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("inventories").
		UpdateOne(ctx, filter, bson.M{"$set": editedInventory})
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("inventories").
		DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Collection("inventories").
		CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
