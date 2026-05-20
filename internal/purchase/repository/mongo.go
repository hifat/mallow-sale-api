package purchaseRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	purchaseModule "github.com/hifat/mallow-sale-api/internal/purchase"
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
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

func NewMongo(db *mongo.Database) purchaseModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *purchaseModule.CreatePurchaseRequest) (string, error) {
	entity := &purchaseModule.Entity{
		PurchaseStatusCode: purchaseStatusModule.EnumPurchaseStatusCodePending,
		Base: utilsModule.Base{
			ID:        database.MustObjectIDFromHex(req.ID),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	res, err := r.db.Collection("purchases").InsertOne(ctx, entity)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*purchaseModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var entity purchaseModule.Entity
	err := r.db.Collection("purchases").FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}
		return nil, err
	}

	return &purchaseModule.Response{
		ID:                 entity.ID.Hex(),
		PurchaseStatusCode: entity.PurchaseStatusCode,
		CreatedAt:          entity.CreatedAt,
		UpdatedAt:          entity.UpdatedAt,
	}, nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("purchases").DeleteOne(ctx, filter)
	return err
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *purchaseModule.CreatePurchaseRequest) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	entity := &purchaseModule.Entity{
		Base: utilsModule.Base{
			UpdatedAt: time.Now(),
		},
	}
	_, err := r.db.Collection("purchases").UpdateOne(ctx, filter, bson.M{"$set": entity})
	return err
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	return r.db.Collection("purchases").CountDocuments(ctx, bson.M{})
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]purchaseModule.Response, error) {
	pipeline := mongo.Pipeline{}

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

	if len(pipeline) == 0 {
		pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.M{}}})
	}

	cursor, err := r.db.Collection("purchases").Aggregate(ctx, pipeline, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var entities []purchaseModule.Entity
	for cursor.Next(ctx) {
		var entity purchaseModule.Entity
		if err := cursor.Decode(&entity); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	res := make([]purchaseModule.Response, 0, len(entities))
	for _, entity := range entities {
		res = append(res, purchaseModule.Response{
			ID:                 entity.ID.Hex(),
			PurchaseStatusCode: entity.PurchaseStatusCode,
			CreatedAt:          entity.CreatedAt,
			UpdatedAt:          entity.UpdatedAt,
		})
	}
	return res, nil
}
