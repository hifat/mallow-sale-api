//go:build wireinject
// +build wireinject

package supplierInventoryDi

import (
	"github.com/google/wire"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	supplierInventoryHandler "github.com/hifat/mallow-sale-api/internal/supplier/inventory/handler"
	supplierInventoryService "github.com/hifat/mallow-sale-api/internal/supplier/inventory/service"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *supplierInventoryHandler.Handler {
	wire.Build(
		// Repository
		supplierRepository.NewMongo,
		inventoryRepository.NewMongo,

		// Service
		logger.New,
		supplierInventoryService.New,

		// Handler
		supplierInventoryHandler.NewRest,
		supplierInventoryHandler.New,
	)

	return &supplierInventoryHandler.Handler{}
}
