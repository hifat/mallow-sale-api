//go:build wireinject
// +build wireinject

package recipeDI

import (
	"github.com/google/wire"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeHandler"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeRepository"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeService"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	recipeRepository.NewMongo,
)

var ServiceSet = wire.NewSet(
	logger.New,
	recipeService.New,
)

var HandlerSet = wire.NewSet(
	helper.New,
	recipeHandler.New,
	recipeHandler.NewRest,
)

func Init(db *mongo.Database, log *zap.Logger) recipeHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)

	return recipeHandler.Handler{}
}
