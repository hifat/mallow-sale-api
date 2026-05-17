package utilsRepository

import (
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoRepository struct{}

func NewMongo() utilsModule.IRepository {
	return &mongoRepository{}
}

func (r *mongoRepository) NewID() string {
	return primitive.NewObjectID().Hex()
}
