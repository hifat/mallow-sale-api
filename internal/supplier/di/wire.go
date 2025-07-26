//go:build wireinject
// +build wireinject

package supplierDi

import (
	"github.com/google/wire"
	supplierHandler "github.com/hifat/mallow-sale-api/internal/supplier/handler"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	supplierService "github.com/hifat/mallow-sale-api/internal/supplier/service"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *supplierHandler.Handler {
	wire.Build(
		// Repository
		supplierRepository.NewMongo,

		// Service
		logger.New,
		supplierService.New,

		// Handler
		supplierHandler.NewRest,
		supplierHandler.New,
	)

	return &supplierHandler.Handler{}
}
