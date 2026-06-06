package purchasePresetRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	purchasePresetModule "github.com/hifat/mallow-sale-api/internal/purchase/preset"
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

func NewMongo(db *mongo.Database) purchasePresetModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *purchasePresetModule.Request) (string, error) {
	newPreset := &purchasePresetModule.Entity{
		SupplierID:   database.MustObjectIDFromHex(req.SupplierID),
		SupplierName: req.SupplierName,
		Inventories:  make([]purchasePresetModule.InventoryEntity, len(req.Inventories)),
	}

	for i, inv := range req.Inventories {
		newPreset.Inventories[i] = purchasePresetModule.InventoryEntity{
			No:               inv.No,
			Name:             inv.Name,
			PurchaseUnitCode: inv.PurchaseUnitCode,
		}
	}

	result, err := r.db.Collection("purchase_presets").InsertOne(ctx, newPreset)
	if err != nil {
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]purchasePresetModule.Response, error) {
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

	cur, err := r.db.Collection("purchase_presets").
		Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	purchasePresets := []purchasePresetModule.Response{}
	for cur.Next(ctx) {
		var preset purchasePresetModule.Entity
		if err := cur.Decode(&preset); err != nil {
			return nil, err
		}
		purchasePresets = append(purchasePresets, purchasePresetModule.Response{
			ID:           preset.ID.Hex(),
			SupplierID:   preset.SupplierID.Hex(),
			SupplierName: preset.SupplierName,
			CreatedAt:    &preset.CreatedAt,
			UpdatedAt:    &preset.UpdatedAt,
			Inventories: func() []purchasePresetModule.InventoryResponse {
				invResponses := make([]purchasePresetModule.InventoryResponse, len(preset.Inventories))
				for i, inv := range preset.Inventories {
					invResponses[i] = purchasePresetModule.InventoryResponse{
						ID:               inv.ID,
						No:               inv.No,
						Name:             inv.Name,
						PurchaseUnitCode: inv.PurchaseUnitCode,
					}
				}
				return invResponses
			}(),
		})
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return purchasePresets, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*purchasePresetModule.Response, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var preset purchasePresetModule.Entity
	err = r.db.Collection("purchase_presets").
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&preset)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	response := &purchasePresetModule.Response{
		ID:           preset.ID.Hex(),
		SupplierID:   preset.SupplierID.Hex(),
		SupplierName: preset.SupplierName,
		CreatedAt:    &preset.CreatedAt,
		UpdatedAt:    &preset.UpdatedAt,
		Inventories: func() []purchasePresetModule.InventoryResponse {
			invResponses := make([]purchasePresetModule.InventoryResponse, len(preset.Inventories))
			for i, inv := range preset.Inventories {
				invResponses[i] = purchasePresetModule.InventoryResponse{
					ID:               inv.ID,
					No:               inv.No,
					Name:             inv.Name,
					PurchaseUnitCode: inv.PurchaseUnitCode,
				}
			}
			return invResponses
		}(),
	}

	return response, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *purchasePresetModule.Request) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"supplier_id":   database.MustObjectIDFromHex(req.SupplierID),
			"supplier_name": req.SupplierName,
			"inventories": func() []purchasePresetModule.InventoryEntity {
				invEntities := make([]purchasePresetModule.InventoryEntity, len(req.Inventories))
				for i, inv := range req.Inventories {
					invEntities[i] = purchasePresetModule.InventoryEntity{
						No:               inv.No,
						Name:             inv.Name,
						PurchaseUnitCode: inv.PurchaseUnitCode,
					}
				}
				return invEntities
			}(),
			"updated_at": time.Now(),
		},
	}

	result, err := r.db.Collection("purchase_presets").
		UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return define.ErrRecordNotFound
	}

	return nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := r.db.Collection("purchase_presets").
		DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return define.ErrRecordNotFound
	}

	return nil
}

func (r *mongoRepository) Count(ctx context.Context, query *utilsModule.QueryReq) (int64, error) {
	filter := bson.M{}
	if query.Search != "" {
		filter["name"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	count, err := r.db.Collection("purchase_presets").CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
