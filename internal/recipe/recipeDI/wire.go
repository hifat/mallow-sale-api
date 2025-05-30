//go:build wireinject
// +build wireinject

package recipeDI

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository"
	"github.com/hifat/mallow-sale-api/internal/recipe/recipeHandler"
	"github.com/hifat/mallow-sale-api/internal/recipe/recipeRepository"
	"github.com/hifat/mallow-sale-api/internal/recipe/recipeService"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	recipeRepository.NewMongo,
	usageUnitRepository.NewMongo,
	inventoryRepository.NewMongo,
	inventoryRepository.NewGRPC,
	usageUnitRepository.NewGRPC,
)

var ServiceSet = wire.NewSet(
	logger.New,
	rules.New,
	helper.New,
	recipeService.New,
)

var HandlerSet = wire.NewSet(
	recipeHandler.New,
	recipeHandler.NewRest,
)

func Init(db *mongo.Database, log *zap.Logger, validate *validator.Validate, grpc rpc.GrpcClient) recipeHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)
	return recipeHandler.Handler{}
}
