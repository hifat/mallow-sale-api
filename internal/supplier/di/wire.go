package supplierDi

import (
	supplierHandler "github.com/hifat/mallow-sale-api/internal/supplier/handler"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	supplierService "github.com/hifat/mallow-sale-api/internal/supplier/service"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Rest *supplierHandler.Rest
}

func Init(db *mongo.Database, logger logger.Logger) *Handler {
	repo := supplierRepository.NewMongo(db)
	svc := supplierService.New(repo, logger)
	rest := supplierHandler.NewRest(svc)

	return &Handler{Rest: rest}
}
