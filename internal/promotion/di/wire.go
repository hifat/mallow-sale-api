//go:build wireinject
// +build wireinject

package promotionDi

import (
	"github.com/google/wire"
	promotionHandler "github.com/hifat/mallow-sale-api/internal/promotion/handler"
	promotionRepository "github.com/hifat/mallow-sale-api/internal/promotion/repository"
	promotionService "github.com/hifat/mallow-sale-api/internal/promotion/service"
	recipeHelper "github.com/hifat/mallow-sale-api/internal/recipe/helper"
	recipeRepository "github.com/hifat/mallow-sale-api/internal/recipe/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *promotionHandler.Handler {
	wire.Build(
		// Repository
		promotionRepository.NewMongo,
		recipeRepository.NewMongo,

		// Helper
		recipeHelper.New,

		// Service
		logger.New,
		promotionService.New,

		// Handler
		promotionHandler.NewRest,
		promotionHandler.New,
	)

	return &promotionHandler.Handler{}
}
