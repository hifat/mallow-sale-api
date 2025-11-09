//go:build wireinject
// +build wireinject

package middlewareDi

import (
	"github.com/google/wire"
	middlewareHandler "github.com/hifat/mallow-sale-api/internal/middleware/handler"
	middlewareService "github.com/hifat/mallow-sale-api/internal/middleware/service"
	userRepository "github.com/hifat/mallow-sale-api/internal/user/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *middlewareHandler.Handler {
	wire.Build(
		// Repository
		userRepository.NewMongo,

		// Service
		logger.New,
		middlewareService.New,

		// Handler
		middlewareHandler.NewRest,
		middlewareHandler.New,
	)

	return &middlewareHandler.Handler{}
}
