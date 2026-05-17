package purchaseSupplierEvidenceRepository

import (
	"context"

	"time"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) purchaseSupplierEvidenceModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *purchaseSupplierEvidenceModule.CreateEvidenceRequest, supplierID string) error {
	entity := &purchaseSupplierEvidenceModule.Entity{
		PurchaseSupplierID:               database.MustObjectIDFromHex(supplierID),
		PurchaseSupplierEvidenceTypeCode: req.Type,
		FileName:                         req.FileName,
		FileRename:                       req.FileRename,
		Path:                             req.Path,
		FileStatusCode:                   req.FileStatusCode,
		UploadedAt:                       req.UploadedAt,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	res, err := r.db.Collection("purchase_supplier_evidences").InsertOne(ctx, entity)
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
	_, err := r.db.Collection("purchase_supplier_evidences").DeleteMany(ctx, filter)
	return err
}

func (r *mongoRepository) FindBySupplierID(ctx context.Context, supplierID string) ([]purchaseSupplierEvidenceModule.Response, error) {
	filter := bson.M{"purchase_supplier_id": database.MustObjectIDFromHex(supplierID)}
	cursor, err := r.db.Collection("purchase_supplier_evidences").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []purchaseSupplierEvidenceModule.Entity
	if err := cursor.All(ctx, &entities); err != nil {
		return nil, err
	}

	res := make([]purchaseSupplierEvidenceModule.Response, 0, len(entities))
	for _, entity := range entities {
		res = append(res, purchaseSupplierEvidenceModule.Response{
			ID:                               entity.ID.Hex(),
			PurchaseSupplierID:               entity.PurchaseSupplierID.Hex(),
			PurchaseSupplierEvidenceTypeCode: entity.PurchaseSupplierEvidenceTypeCode,
			FileName:                         entity.FileName,
			FileRename:                       entity.FileRename,
			Path:                             entity.Path,
			FileStatusCode:                   entity.FileStatusCode,
			UploadedAt:                       entity.UploadedAt,
			CreatedAt:                        entity.CreatedAt,
			UpdatedAt:                        entity.UpdatedAt,
		})
	}
	return res, nil
}
