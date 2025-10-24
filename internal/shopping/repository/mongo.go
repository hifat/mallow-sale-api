package shoppingRepository

import (
	"context"
	"errors"
	"fmt"
	"time"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/database"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IRepository {
	return &mongoRepository{db}
}

func (r *mongoRepository) Find(ctx context.Context) ([]shoppingModule.Response, error) {
	cur, err := r.db.Collection("shoppings").
		Find(ctx, bson.M{}, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	shoppings := []shoppingModule.Response{}
	for cur.Next(ctx) {
		var shopping shoppingModule.Entity
		if err := cur.Decode(&shopping); err != nil {
			return nil, err
		}

		shoppings = append(shoppings, shoppingModule.Response{
			ID:               shopping.ID.Hex(),
			Name:             shopping.Name,
			IsComplete:       shopping.IsComplete,
			PurchaseQuantity: shopping.PurchaseQuantity,
			PurchaseUnit: usageUnitModule.Prototype{
				Code: shopping.PurchaseUnit.Code,
				Name: shopping.PurchaseUnit.Name,
			},
		})
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return shoppings, nil
}
func (r *mongoRepository) FindByID(ctx context.Context, id string) (*shoppingModule.Response, error) {
	fmt.Println(database.MustObjectIDFromHex(id))
	var shopping shoppingModule.Entity
	err := r.db.Collection("shoppings").
		FindOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		}).
		Decode(&shopping)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	res := &shoppingModule.Response{
		ID:               shopping.ID.Hex(),
		Name:             shopping.Name,
		PurchaseQuantity: shopping.PurchaseQuantity,
		PurchaseUnit:     usageUnitModule.Prototype(shopping.PurchaseUnit),
	}

	return res, nil
}

func (r *mongoRepository) Create(ctx context.Context, req *shoppingModule.Request) error {
	newShopping := shoppingModule.Entity{
		Name:             req.Name,
		PurchaseQuantity: req.PurchaseQuantity,
		PurchaseUnit:     usageUnitModule.Entity(req.PurchaseUnit),
		IsComplete:       false,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := r.db.Collection("shoppings").
		InsertOne(ctx, newShopping)

	return err
}

func (r *mongoRepository) UpdateIsComplete(ctx context.Context, id string, req *shoppingModule.ReqUpdateIsComplete) error {
	_, err := r.db.Collection("shoppings").
		UpdateOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		}, bson.M{"$set": bson.M{
			"is_complete": req.IsComplete,
		}})

	return err
}

func (r *mongoRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := r.db.Collection("shoppings").
		DeleteOne(ctx, bson.M{
			"_id": database.MustObjectIDFromHex(id),
		})

	return err
}
