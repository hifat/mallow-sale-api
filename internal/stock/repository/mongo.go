package stockRepository

import (
	"context"
	"strings"
	"time"

	stockModule "github.com/hifat/mallow-sale-api/internal/stock"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) stockModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *stockModule.Request) error {
	newStock := &stockModule.Entity{
		InventoryID:      req.InventoryID,
		SupplierID:       req.SupplierID,
		PurchasePrice:    req.PurchasePrice,
		PurchaseQuantity: req.PurchaseQuantity,
		PurchaseUnit: usageUnitModule.Entity{
			Code: req.PurchaseUnit.Code,
			Name: req.PurchaseUnit.Name,
		},
		Remark: req.Remark,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	_, err := r.db.Collection("stocks").InsertOne(ctx, newStock)
	return err
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]stockModule.Response, error) {
	pipeline := mongo.Pipeline{}
	if query.Search != "" {
		setStage := bson.D{
			{
				Key: "$set", Value: bson.M{
					"inventory_id": bson.M{
						"$toObjectId": "$inventory_id",
					},
				},
			},
		}

		lookupStage := bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "inventories",
			"localField":   "inventory_id",
			"foreignField": "_id",
			"as":           "inventory",
		}}}

		unwindStage := bson.D{{Key: "$unwind", Value: "$inventory"}}

		matchStage := bson.D{{Key: "$match", Value: bson.M{
			"inventory.name": bson.M{
				"$regex":   query.Search,
				"$options": "i",
			},
		}}}

		// 4. Run the pipeline
		pipeline = mongo.Pipeline{setStage, lookupStage, unwindStage, matchStage}
	}

	if query.Sort != "" && query.Order != "" {
		order := 1
		if strings.ToLower(query.Order) == "desc" {
			order = -1
		}
		// เพิ่ม stage $sort
		pipeline = append(pipeline, bson.D{{Key: "$sort", Value: bson.M{query.Sort: order}}})
	}

	if query.Page > 0 && query.Limit > 0 {
		skip := int64((query.Page - 1) * query.Limit)
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: skip}})
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: int64(query.Limit)}})
	}

	if query.Fields != "" {
		fields := strings.Split(query.Fields, ",")
		projection := bson.M{}
		for _, field := range fields {
			projection[field] = 1
		}
		pipeline = append(pipeline, bson.D{{Key: "$project", Value: projection}})
	}

	// cursor, err := r.db.Collection("stocks").Find(ctx, filter, findOptions)
	cursor, err := r.db.Collection("stocks").Aggregate(ctx, pipeline, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []stockModule.Entity
	for cursor.Next(ctx) {
		var stock stockModule.Entity
		if err := cursor.Decode(&stock); err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	res := make([]stockModule.Response, 0, len(stocks))
	for _, stock := range stocks {
		createdAt := stock.CreatedAt
		updatedAt := stock.UpdatedAt
		res = append(res, stockModule.Response{
			Prototype: stockModule.Prototype{
				ID:               stock.ID.Hex(),
				InventoryID:      stock.InventoryID,
				Inventory:        nil, // Will be populated by service layer if needed
				SupplierID:       stock.SupplierID,
				Supplier:         nil, // Will be populated by service layer if needed
				PurchasePrice:    stock.PurchasePrice,
				PurchaseQuantity: stock.PurchaseQuantity,
				PurchaseUnit: usageUnitModule.Prototype{
					Code: stock.PurchaseUnit.Code,
					Name: stock.PurchaseUnit.Name,
				},
				Remark:    stock.Remark,
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
		})
	}
	return res, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*stockModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var stock stockModule.Entity
	if err := r.db.Collection("stocks").FindOne(ctx, filter).Decode(&stock); err != nil {
		return nil, err
	}
	createdAt := stock.CreatedAt
	updatedAt := stock.UpdatedAt
	return &stockModule.Response{
		Prototype: stockModule.Prototype{
			ID:               stock.ID.Hex(),
			InventoryID:      stock.InventoryID,
			Inventory:        nil, // Will be populated by service layer if needed
			SupplierID:       stock.SupplierID,
			Supplier:         nil, // Will be populated by service layer if needed
			PurchasePrice:    stock.PurchasePrice,
			PurchaseQuantity: stock.PurchaseQuantity,
			PurchaseUnit: usageUnitModule.Prototype{
				Code: stock.PurchaseUnit.Code,
				Name: stock.PurchaseUnit.Name,
			},
			Remark:    stock.Remark,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		},
	}, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *stockModule.Request) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	update := bson.M{"$set": bson.M{
		"inventory_id":      req.InventoryID,
		"supplier_id":       req.SupplierID,
		"purchase_price":    req.PurchasePrice,
		"purchase_quantity": req.PurchaseQuantity,
		"purchase_unit": bson.M{
			"code": req.PurchaseUnit.Code,
			"name": req.PurchaseUnit.Name,
		},
		"remark":     req.Remark,
		"updated_at": time.Now(),
	}}
	_, err := r.db.Collection("stocks").UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("stocks").DeleteOne(ctx, filter)
	return err
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	return r.db.Collection("stocks").CountDocuments(ctx, bson.M{})
}
