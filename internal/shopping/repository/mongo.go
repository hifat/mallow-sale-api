package shoppingRepository

import (
	"context"
	"errors"
	"time"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) shoppingModule.IRepository {
	return &mongoRepository{db}
}

func (r *mongoRepository) Find(ctx context.Context) ([]shoppingModule.Response, error) {
	cur, err := r.db.Collection("shoppings").
		Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	shoppings := []shoppingModule.Response{}
	for cur.Next(ctx) {
		var shopping shoppingModule.Entity
		if err := cur.Decode(&shopping); err != nil {
			return nil, err
		}

		shoppings = append(shoppings, shoppingModule.Response{
			ID:           shopping.ID.Hex(),
			SupplierID:   shopping.SupplierID,
			SupplierName: shopping.SupplierName,
			Inventories:  make([]shoppingModule.PrototypeInventory, 0),
		})
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return shoppings, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*shoppingModule.Response, error) {
	var shopping shoppingModule.Entity
	err := r.db.Collection("shoppings").
		FindOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		}).
		Decode(&shopping)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	res := &shoppingModule.Response{
		ID:           shopping.ID.Hex(),
		SupplierID:   shopping.SupplierID,
		SupplierName: shopping.SupplierName,
		Inventories:  make([]shoppingModule.PrototypeInventory, 0, len(shopping.Inventories)),
	}

	for _, v := range shopping.Inventories {
		res.Inventories = append(res.Inventories, shoppingModule.PrototypeInventory(v))
	}

	return res, nil
}

func (r *mongoRepository) Create(ctx context.Context, req *shoppingModule.Request) error {
	newShopping := shoppingModule.Entity{
		SupplierID:   req.SupplierID,
		SupplierName: req.SupplierName,
		Inventories:  make([]shoppingModule.Inventory, 0, len(req.Inventories)),
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, v := range req.Inventories {
		newShopping.Inventories = append(newShopping.Inventories, shoppingModule.Inventory{
			OrderNo:          v.OrderNo,
			InventoryID:      v.InventoryID,
			InventoryName:    v.InventoryName,
			PurchaseUnit:     usageUnitModule.Entity(v.PurchaseUnit),
			PurchaseQuantity: v.PurchaseQuantity,
			Status:           shoppingModule.InventoryStatus(v.Status),
		})
	}

	_, err := r.db.Collection("shoppings").
		InsertOne(ctx, newShopping)

	return err
}

func (r *mongoRepository) UpdateIsComplete(ctx context.Context, id string, req *shoppingModule.ReqUpdateIsComplete) error {
	_, err := r.db.Collection("shoppings").
		UpdateOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		}, bson.M{"$set": bson.M{
			"is_complete": req.IsComplete,
		}})

	return err
}

func (r *mongoRepository) ReOrderNo(ctx context.Context, reqs []shoppingModule.ReqReOrder) error {
	models := []mongo.WriteModel{}

	for _, v := range reqs {
		update := mongo.NewUpdateOneModel().
			SetFilter(bson.M{"_id": database.MustObjectIDFromHex(v.ID)}).
			SetUpdate(bson.M{
				"$set": bson.M{
					"order_no":   v.OrderNo,
					"updated_at": time.Now(),
				},
			})
		models = append(models, update)
	}

	if len(models) > 0 {
		_, err := r.db.Collection("shoppings").
			BulkWrite(ctx, models)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.Collection("shoppings").
		DeleteOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		})

	return err
}
