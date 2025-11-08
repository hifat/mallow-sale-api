//go:build wireinject
// +build wireinject

package authDi

import (
	"github.com/google/wire"
	authHandler "github.com/hifat/mallow-sale-api/internal/auth/handler"
	authService "github.com/hifat/mallow-sale-api/internal/auth/service"
	userRepository "github.com/hifat/mallow-sale-api/internal/user/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *authHandler.Handler {
	wire.Build(
		// Repository
		userRepository.NewMongo,

		// Service
		logger.New,
		authService.New,

		// Handler
		authHandler.NewRest,
		authHandler.New,
	)

	return &authHandler.Handler{}
}
