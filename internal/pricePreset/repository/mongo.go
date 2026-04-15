package pricePresetRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	pricePresetModule "github.com/hifat/mallow-sale-api/internal/pricePreset"
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

func NewMongo(db *mongo.Database) pricePresetModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *pricePresetModule.Request) error {
	newPreset := &pricePresetModule.Entity{
		InventoryID: req.InventoryID,
		Prices: []pricePresetModule.Price{
			{
				ID:        primitive.NewObjectID().Hex(),
				StockID:   req.StockID,
				Price:     req.Price,
				CreatedAt: time.Now(),
			},
		},
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	_, err := r.db.Collection("price_presets").InsertOne(ctx, newPreset)
	return err
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]pricePresetModule.Response, error) {
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

		pipeline = mongo.Pipeline{setStage, lookupStage, unwindStage, matchStage}
	}

	if query.Sort != "" && query.Order != "" {
		order := 1
		if strings.ToLower(query.Order) == "desc" {
			order = -1
		}
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

	cursor, err := r.db.Collection("price_presets").Aggregate(ctx, pipeline, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var pricePresets []pricePresetModule.Entity
	if err = cursor.All(ctx, &pricePresets); err != nil {
		return nil, err
	}

	res := make([]pricePresetModule.Response, 0, len(pricePresets))
	for _, preset := range pricePresets {
		createdAt := preset.CreatedAt
		updatedAt := preset.UpdatedAt

		prices := make([]pricePresetModule.PricePrototype, 0, len(preset.Prices))
		for _, p := range preset.Prices {
			prices = append(prices, pricePresetModule.PricePrototype{
				ID:        p.ID,
				StockID:   p.StockID,
				Price:     p.Price,
				CreatedAt: p.CreatedAt,
			})
		}

		res = append(res, pricePresetModule.Response{
			Prototype: pricePresetModule.Prototype{
				ID:          preset.ID.Hex(),
				InventoryID: preset.InventoryID,
				Inventory:   nil, // Will be populated by service layer if needed
				Prices:      prices,
				CreatedAt:   &createdAt,
				UpdatedAt:   &updatedAt,
			},
		})
	}
	return res, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*pricePresetModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var preset pricePresetModule.Entity
	if err := r.db.Collection("price_presets").FindOne(ctx, filter).Decode(&preset); err != nil {
		return nil, err
	}

	createdAt := preset.CreatedAt
	updatedAt := preset.UpdatedAt

	prices := make([]pricePresetModule.PricePrototype, 0, len(preset.Prices))
	for _, p := range preset.Prices {
		prices = append(prices, pricePresetModule.PricePrototype{
			ID:        p.ID,
			StockID:   p.StockID,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
		})
	}

	return &pricePresetModule.Response{
		Prototype: pricePresetModule.Prototype{
			ID:          preset.ID.Hex(),
			InventoryID: preset.InventoryID,
			Prices:      prices,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		},
	}, nil
}

func (r *mongoRepository) FindByInventoryID(ctx context.Context, inventoryID string) (*pricePresetModule.Entity, error) {
	filter := bson.M{"inventory_id": inventoryID}
	var preset pricePresetModule.Entity
	err := r.db.Collection("price_presets").FindOne(ctx, filter).Decode(&preset)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}
		return nil, err
	}
	return &preset, nil
}

func (r *mongoRepository) FindByPriceID(ctx context.Context, priceID string) (*pricePresetModule.Entity, error) {
	filter := bson.M{"prices.id": priceID}
	var preset pricePresetModule.Entity
	err := r.db.Collection("price_presets").FindOne(ctx, filter).Decode(&preset)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}
		return nil, err
	}
	return &preset, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *pricePresetModule.Request) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	update := bson.M{
		"$push": bson.M{
			"prices": pricePresetModule.Price{
				ID:        primitive.NewObjectID().Hex(),
				StockID:   req.StockID,
				Price:     req.Price,
				CreatedAt: time.Now(),
			},
		},
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}
	_, err := r.db.Collection("price_presets").UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("price_presets").DeleteOne(ctx, filter)
	return err
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	return r.db.Collection("price_presets").CountDocuments(ctx, bson.M{})
}
