package storagerepo

import (
	"context"
	"time"

	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) storageModule.IRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Create(ctx context.Context, req *storageModule.CreateStorageRequest) (*storageModule.UploadResponse, error) {
	entity := &storageModule.Entity{
		FileName:   req.Filename,
		ObjectKey:  req.ObjectKey,
		StatusCode: req.StatusCode,
		Base: utilsModule.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := r.db.Collection("storages").InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &storageModule.UploadResponse{
		Filename:  req.Filename,
		ObjectKey: req.ObjectKey,
	}, nil
}
