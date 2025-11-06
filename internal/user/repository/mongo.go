package userRepository

import (
	"context"
	"errors"

	userModule "github.com/hifat/mallow-sale-api/internal/user"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) IRepository {
	return &mongoRepository{
		db: db,
	}
}

func (r *mongoRepository) FindByUsername(ctx context.Context, username string) (*userModule.Response, error) {
	filter := bson.M{"username": username}

	var user userModule.Entity
	err := r.db.Collection("users").
		FindOne(ctx, filter).
		Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, define.ErrRecordNotFound
		}

		return nil, err
	}

	res := &userModule.Response{
		Prototype: userModule.Prototype{
			Name:     user.Name,
			Username: user.Username,
			Password: user.Password,
		},
	}

	return res, nil
}
