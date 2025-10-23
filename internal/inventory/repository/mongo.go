package inventoryRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IRepository {
	return &mongoRepository{
		db: db,
	}
}

func (r *mongoRepository) Create(ctx context.Context, req *inventoryModule.Request) error {
	newInventory := &inventoryModule.Entity{
		Name: req.Name,
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

	return err
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

func (r *mongoRepository) FindByName(ctx context.Context, name string) (*inventoryModule.Response, error) {
	filter := bson.M{"name": name}
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

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]inventoryModule.Response, error) {
	filter := bson.M{}
	if query.Search != "" {
		filter["name"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	findOptions := options.Find()

	if query.Sort != "" && query.Order != "" {
		order := 1
		if strings.ToLower(query.Order) == "desc" {
			order = -1
		}
		findOptions.SetSort(bson.M{query.Sort: order})
	}

	if query.Page > 0 && query.Limit > 0 {
		findOptions.SetSkip(int64((query.Page - 1) * query.Limit))
		findOptions.SetLimit(int64(query.Limit))
	}

	if query.Fields != "" {
		fields := strings.Split(query.Fields, ",")
		projection := bson.M{}
		for _, field := range fields {
			projection[field] = 1
		}

		findOptions.SetProjection(projection)
	}

	cur, err := r.db.Collection("inventories").
		Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var inventories []inventoryModule.Entity
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
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("inventories").
		UpdateOne(ctx, filter, bson.M{"$set": bson.M{
			"name":             req.Name,
			"purchase_unit":    req.PurchaseUnit,
			"yield_percentage": req.YieldPercentage,
			"remark":           req.Remark,
			"updated_at":       time.Now(),
		}})
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

func (r *mongoRepository) UpdateStock(ctx context.Context, id string, currentQuantity float64, purchasePrice float64) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	update := bson.M{
		"$set": bson.M{
			"purchase_quantity": currentQuantity,
			"purchase_price":    purchasePrice,
			"updated_at":        time.Now(),
		},
	}
	_, err := r.db.Collection("inventories").UpdateOne(ctx, filter, update)
	return err
}
