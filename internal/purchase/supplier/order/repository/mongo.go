package purchaseSupplierOrderRepository

import (
	"context"
	"time"

	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) purchaseSupplierOrderModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *purchaseSupplierOrderModule.CreateOrderRequest, supplierID string) error {
	entity := &purchaseSupplierOrderModule.Entity{
		PurchaseSupplierID: database.MustObjectIDFromHex(supplierID),
		InventoryID:        database.MustObjectIDFromHex(req.InventoryID),
		InventoryName:      req.InventoryName,
		Quantity:           req.Quantity,
		UsageUnitCode:      req.UsageUnitCode,
		StatusCode:         req.StatusCode,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	res, err := r.db.Collection("purchase_supplier_orders").InsertOne(ctx, entity)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		entity.ID = oid
	}
	return nil
}

func (r *mongoRepository) DeleteBySupplierID(ctx context.Context, supplierID string) error {
	filter := bson.M{"purchase_supplier_id": database.MustObjectIDFromHex(supplierID)}
	_, err := r.db.Collection("purchase_supplier_orders").DeleteMany(ctx, filter)
	return err
}

func (r *mongoRepository) FindBySupplierID(ctx context.Context, supplierID string) ([]purchaseSupplierOrderModule.Response, error) {
	filter := bson.M{"purchase_supplier_id": database.MustObjectIDFromHex(supplierID)}
	cursor, err := r.db.Collection("purchase_supplier_orders").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []purchaseSupplierOrderModule.Entity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	res := make([]purchaseSupplierOrderModule.Response, 0, len(entities))
	for _, entity := range entities {
		res = append(res, purchaseSupplierOrderModule.Response{
			ID:                 entity.ID.Hex(),
			PurchaseSupplierID: entity.PurchaseSupplierID.Hex(),
			InventoryID:        entity.InventoryID.Hex(),
			InventoryName:      entity.InventoryName,
			Quantity:           entity.Quantity,
			UsageUnitCode:      entity.UsageUnitCode,
			StatusCode:         entity.StatusCode,
			CreatedAt:          entity.CreatedAt,
			UpdatedAt:          entity.UpdatedAt,
		})
	}
	return res, nil
}
