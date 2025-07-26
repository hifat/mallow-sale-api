//go:build wireinject
// +build wireinject

package stockDi

import (
	"github.com/google/wire"
	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	stockHandler "github.com/hifat/mallow-sale-api/internal/stock/handler"
	stockRepository "github.com/hifat/mallow-sale-api/internal/stock/repository"
	stockService "github.com/hifat/mallow-sale-api/internal/stock/service"
	supplierHelper "github.com/hifat/mallow-sale-api/internal/supplier/helper"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *stockHandler.Handler {
	wire.Build(
		// Repositories
		stockRepository.NewMongo,
		inventoryRepository.NewMongo,
		supplierRepository.NewMongo,
		usageUnitRepository.NewMongo,

		// Helpers
		inventoryHelper.New,
		supplierHelper.New,
		usageUnitHelper.New,

		// Service
		logger.New,
		stockService.New,

		// Handler
		stockHandler.NewRest,
		stockHandler.New,
	)

	return &stockHandler.Handler{}
}
