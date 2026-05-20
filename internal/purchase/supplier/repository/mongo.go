package purchaseSupplierRepository

import (
	"context"
	"time"

	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) purchaseSupplierModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *purchaseSupplierModule.CreateSupplierRequest, purchaseID string) (string, error) {
	entity := &purchaseSupplierModule.Entity{
		ID:              database.MustObjectIDFromHex(req.PurchaseSupplierID),
		PurchaseID:      database.MustObjectIDFromHex(purchaseID),
		SupplierID:      database.MustObjectIDFromHex(req.SupplierID),
		SupplierName:    req.SupplierName,
		StatusCode:      req.StatusCode,
		PaymentTypeCode: req.PaymentType,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	res, err := r.db.Collection("purchase_suppliers").InsertOne(ctx, entity)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *mongoRepository) DeleteByPurchaseID(ctx context.Context, purchaseID string) error {
	filter := bson.M{"purchase_id": database.MustObjectIDFromHex(purchaseID)}
	_, err := r.db.Collection("purchase_suppliers").DeleteMany(ctx, filter)

	return err
}

func (r *mongoRepository) FindByPurchaseID(ctx context.Context, purchaseID string) ([]purchaseSupplierModule.Response, error) {
	filter := bson.M{"purchase_id": database.MustObjectIDFromHex(purchaseID)}
	cursor, err := r.db.Collection("purchase_suppliers").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []purchaseSupplierModule.Entity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	res := make([]purchaseSupplierModule.Response, 0, len(entities))
	for _, entity := range entities {
		res = append(res, purchaseSupplierModule.Response{
			ID:              entity.ID.Hex(),
			PurchaseID:      entity.PurchaseID.Hex(),
			SupplierID:      entity.SupplierID.Hex(),
			SupplierName:    entity.SupplierName,
			StatusCode:      entity.StatusCode,
			PaymentTypeCode: entity.PaymentTypeCode,
			CreatedAt:       entity.CreatedAt,
			UpdatedAt:       entity.UpdatedAt,
		})
	}
	return res, nil
}
