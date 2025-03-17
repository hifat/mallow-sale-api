//go:build wireinject
// +build wireinject

package inventoryDI

import (
	"github.com/google/wire"
	"github.com/hifat/cost-calculator-api/config"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryHandler"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryRepository"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryService"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	inventoryRepository.NewMongo,
)

var ServiceSet = wire.NewSet(
	helper.New,
	logger.New,
	inventoryService.New,
)

var HandlerSet = wire.NewSet(
	inventoryHandler.New,
	inventoryHandler.NewRest,
)

func InitInventory(cfg *config.Config, db *mongo.Database, log *zap.Logger) inventoryHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)

	return inventoryHandler.Handler{}
}
