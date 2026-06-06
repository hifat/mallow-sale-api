//go:build wireinject
// +build wireinject

package purchaseDi

import (
	"github.com/google/wire"
	purchasePresetHandler "github.com/hifat/mallow-sale-api/internal/purchase/preset/handler"
	purchasePresetRepository "github.com/hifat/mallow-sale-api/internal/purchase/preset/repository"
	purchaseService "github.com/hifat/mallow-sale-api/internal/purchase/preset/service"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *purchasePresetHandler.Handler {
	wire.Build(
		// Repository
		purchasePresetRepository.NewMongo,

		// Service
		logger.New,
		purchaseService.New,

		// Handler
		purchasePresetHandler.NewRest,
		purchasePresetHandler.New,
	)

	return &purchasePresetHandler.Handler{}
}
