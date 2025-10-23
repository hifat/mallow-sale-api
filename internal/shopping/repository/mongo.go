package shoppingRepository

import (
	"context"
	"time"

	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) IRepository {
	return &mongoRepository{db}
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

func (r *mongoRepository) UpdateIsComplete(ctx context.Context, req *shoppingModule.UpdateIsComplete) error {
	return nil
}

func (r *mongoRepository) Delete(ctx context.Context, id string) error {
	return nil
}
