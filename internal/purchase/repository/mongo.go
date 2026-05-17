package purchaseRepository

import (
	"context"
	"errors"
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
