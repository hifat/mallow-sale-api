//go:build wireinject
// +build wireinject

package mockInventoryDI

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryHandler"
	mockInventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository/mock"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryService"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	mockInventoryRepository.NewMockIInventoryGRPCRepository,
	mockInventoryRepository.NewMockIInventoryRepository,
)

var ServiceSet = wire.NewSet(
	helper.New,
	logger.New,
	rules.New,
	inventoryService.New,
	inventoryService.NewGRPC,
)

var HandlerSet = wire.NewSet(
	inventoryHandler.New,
	inventoryHandler.NewRest,
	inventoryHandler.NewGRPC,
)

func Init(cfg *config.Config, db *mongo.Database, log *zap.Logger, validator *validator.Validate, grpc rpc.GrpcClient) inventoryHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)

	return inventoryHandler.Handler{}
}
