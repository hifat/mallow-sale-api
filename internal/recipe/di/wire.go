//go:build wireinject
// +build wireinject

package recipeDi

import (
	"github.com/google/wire"
	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	recipeHandler "github.com/hifat/mallow-sale-api/internal/recipe/handler"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
	recipeService "github.com/hifat/mallow-sale-api/internal/recipe/service"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *recipeHandler.Handler {
	wire.Build(
		// Repository
		recipeRepository.NewMongo,
		inventoryRepository.NewMongo,
		usageUnitRepository.NewMongo,

		// Helper
		usageUnitHelper.New,
		inventoryHelper.New,

		// Service
		logger.New,
		recipeService.New,

		// Handler
		recipeHandler.NewRest,
		recipeHandler.New,
	)

	return &recipeHandler.Handler{}
}
