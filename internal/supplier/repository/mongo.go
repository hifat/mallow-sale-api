package supplierRepository

import (
	"context"
	"strings"
	"time"

	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *supplierModule.Request) error {
	newSupplier := &supplierModule.Entity{
		Name:   req.Name,
		ImgUrl: req.ImgUrl,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	_, err := r.db.Collection("suppliers").InsertOne(ctx, newSupplier)
	return err
}

func (r *mongoRepository) Find(ctx context.Context, query *utilsModule.QueryReq) ([]supplierModule.Response, error) {
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

	cursor, err := r.db.Collection("suppliers").Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var suppliers []supplierModule.Entity
	for cursor.Next(ctx) {
		var supplier supplierModule.Entity
		if err := cursor.Decode(&supplier); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}

	res := make([]supplierModule.Response, 0, len(suppliers))
	for _, supplier := range suppliers {
		createdAt := supplier.CreatedAt
		updatedAt := supplier.UpdatedAt
		res = append(res, supplierModule.Response{
			Prototype: supplierModule.Prototype{
				ID:        supplier.ID.Hex(),
				Name:      supplier.Name,
				ImgUrl:    supplier.ImgUrl,
				CreatedAt: &createdAt,
				UpdatedAt: &updatedAt,
			},
		})
	}
	return res, nil
}

func (r *mongoRepository) FindByID(ctx context.Context, id string) (*supplierModule.Response, error) {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	var supplier supplierModule.Entity
	if err := r.db.Collection("suppliers").FindOne(ctx, filter).Decode(&supplier); err != nil {
		return nil, err
	}
	createdAt := supplier.CreatedAt
	updatedAt := supplier.UpdatedAt
	return &supplierModule.Response{
		Prototype: supplierModule.Prototype{
			ID:        supplier.ID.Hex(),
			Name:      supplier.Name,
			ImgUrl:    supplier.ImgUrl,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		},
	}, nil
}

func (r *mongoRepository) UpdateByID(ctx context.Context, id string, req *supplierModule.Request) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	update := bson.M{"$set": bson.M{
		"name":       req.Name,
		"img_url":    req.ImgUrl,
		"updated_at": time.Now(),
	}}
	_, err := r.db.Collection("suppliers").UpdateOne(ctx, filter, update)
	return err
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": database.MustObjectIDFromHex(id)}
	_, err := r.db.Collection("suppliers").DeleteOne(ctx, filter)
	return err
}

func (r *mongoRepository) Count(ctx context.Context) (int64, error) {
	return r.db.Collection("suppliers").CountDocuments(ctx, bson.M{})
}
