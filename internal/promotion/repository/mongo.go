package promotionRepository

import (
	"context"
	"errors"
	"strings"
	"time"

	promotionModule "github.com/hifat/mallow-sale-api/internal/promotion"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *promotionModule.Request) error {
	newPromotion := promotionModule.Entity{
		Type: promotionModule.PromotionType{
			Code: req.Type.Code,
			Name: req.Type.Name,
		},
		Name:     req.Name,
		Detail:   req.Detail,
		Discount: req.Discount,
		Price:    req.Price,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := r.db.Collection("promotions").InsertOne(ctx, newPromotion)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]promotionModule.Response, error) {
	filter := bson.M{}
	if query.Search != "" {
		filter["name"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	findOptions := options.Find()
	if query.Limit > 0 {
		findOptions.SetLimit(int64(query.Limit))
	}
	if query.Page > 1 {
		findOptions.SetSkip(int64((query.Page - 1) * query.Limit))
	}
	if query.Sort != "" {
		order := 1
		if strings.ToLower(query.Order) == "desc" {
			order = -1
		}
		findOptions.SetSort(bson.M{query.Sort: order})
	} else {
		findOptions.SetSort(bson.M{"created_at": -1})
	}

	if query.Fields != "" {
		fields := strings.Split(query.Fields, ",")
		projection := bson.M{}
		for _, field := range fields {
			projection[field] = 1
		}
		findOptions.SetProjection(projection)
	}

	promotions := make([]promotionModule.Entity, 0)
	cursor, err := r.db.Collection("promotions").Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var promotion promotionModule.Entity
		if err := cursor.Decode(&promotion); err != nil {
			return nil, err
		}
		promotions = append(promotions, promotion)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	responses := make([]promotionModule.Response, 0, len(promotions))
	for _, promotion := range promotions {
		responses = append(responses, promotionModule.Response{
			ProtoType: promotionModule.ProtoType{
				ID: promotion.ID.Hex(),
				Type: promotionModule.PromotionTypeResponse{
					ID:   promotion.Type.ID.Hex(),
					Code: promotion.Type.Code,
					Name: promotion.Type.Name,
				},
				Name:      promotion.Name,
				Detail:    promotion.Detail,
				Discount:  promotion.Discount,
				Price:     promotion.Price,
				CreatedAt: promotion.CreatedAt.Format(time.RFC3339),
				UpdatedAt: promotion.UpdatedAt.Format(time.RFC3339),
			},
		})
	}

	return responses, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*promotionModule.Response, error) {
	objectID := database.MustObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}
	var promotion promotionModule.Entity

	err := r.db.Collection("promotions").FindOne(ctx, filter).Decode(&promotion)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	response := &promotionModule.Response{
		ProtoType: promotionModule.ProtoType{
			ID: promotion.ID.Hex(),
			Type: promotionModule.PromotionTypeResponse{
				ID:   promotion.Type.ID.Hex(),
				Code: promotion.Type.Code,
				Name: promotion.Type.Name,
			},
			Name:      promotion.Name,
			Detail:    promotion.Detail,
			Discount:  promotion.Discount,
			Price:     promotion.Price,
			CreatedAt: promotion.CreatedAt.Format(time.RFC3339),
			UpdatedAt: promotion.UpdatedAt.Format(time.RFC3339),
		},
	}

	return response, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *promotionModule.Request) error {
	objectID := database.MustObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"type":       req.Type,
			"name":       req.Name,
			"detail":     req.Detail,
			"discount":   req.Discount,
			"price":      req.Price,
			"updated_at": time.Now(),
		},
	}

	_, err := r.db.Collection("promotions").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	objectID := database.MustObjectIDFromHex(id)

	filter := bson.M{"_id": objectID}
	_, err := r.db.Collection("promotions").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	count, err := r.db.Collection("promotions").CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
