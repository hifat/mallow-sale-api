//go:build wireinject
// +build wireinject

package recipeDI

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeHandler"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeRepository"
	"github.com/hifat/cost-calculator-api/internal/recipe/recipeService"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/goroger-core/rules"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	recipeRepository.NewMongo,
)

var ServiceSet = wire.NewSet(
	logger.New,
	rules.New,
	recipeService.New,
)

var HandlerSet = wire.NewSet(
	helper.New,
	recipeHandler.New,
	recipeHandler.NewRest,
)

func Init(db *mongo.Database, log *zap.Logger, validate *validator.Validate) recipeHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)

	return recipeHandler.Handler{}
}
