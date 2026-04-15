//go:build wireinject
// +build wireinject

package pricePresetDi

import (
	"github.com/google/wire"
	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	pricePresetHandler "github.com/hifat/mallow-sale-api/internal/pricePreset/handler"
	pricePresetRepository "github.com/hifat/mallow-sale-api/internal/pricePreset/repository"
	pricePresetService "github.com/hifat/mallow-sale-api/internal/pricePreset/service"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *pricePresetHandler.Handler {
	wire.Build(
		// Repositories
		pricePresetRepository.NewMongo,
		inventoryRepository.NewMongo,

		// Helpers
		inventoryHelper.New,

		// Service
		logger.New,
		pricePresetService.New,

		// Handler
		pricePresetHandler.NewRest,
		pricePresetHandler.New,
	)

	return &pricePresetHandler.Handler{}
}
