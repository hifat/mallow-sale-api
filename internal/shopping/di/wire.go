//go:build wireinject
// +build wireinject

package shoppingDi

import (
	"github.com/google/wire"
	shoppingHandler "github.com/hifat/mallow-sale-api/internal/shopping/handler"
	shoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository"
	shoppingService "github.com/hifat/mallow-sale-api/internal/shopping/service"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *shoppingHandler.Handler {
	wire.Build(
		// Repository
		shoppingRepository.NewMongo,
		usageUnitRepository.NewMongo,

		// Service
		logger.New,
		shoppingService.New,
		shoppingService.NewReceipt,

		// Handler
		shoppingHandler.NewReceiptRest,
		shoppingHandler.NewRest,
		shoppingHandler.New,
	)

	return &shoppingHandler.Handler{}
}
