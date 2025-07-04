//go:build wireinject
// +build wireinject

package inventoryDi

import (
	"github.com/google/wire"
	inventoryHandler "github.com/hifat/mallow-sale-api/internal/inventory/handler"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	inventoryService "github.com/hifat/mallow-sale-api/internal/inventory/service"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *inventoryHandler.Handler {
	wire.Build(
		// Repository
		inventoryRepository.NewMongo,
		usageUnitRepository.NewMongo,

		// Service
		logger.New,
		inventoryService.New,

		// Handler
		inventoryHandler.NewRest,
		inventoryHandler.New,
	)

	return &inventoryHandler.Handler{}
}
