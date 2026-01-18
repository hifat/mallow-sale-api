//go:build wireinject
// +build wireinject

package storageDi

import (
	"github.com/google/wire"
	storageHandler "github.com/hifat/mallow-sale-api/internal/storage/handler"
	storageRepository "github.com/hifat/mallow-sale-api/internal/storage/repository"
	storageService "github.com/hifat/mallow-sale-api/internal/storage/service"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

func Init(cfg *config.Config) (*storageHandler.Handler, error) {
	wire.Build(
		// Repository
		storageRepository.NewGDrive,

		// Service
		logger.New,
		storageService.New,

		// Handler
		storageHandler.NewRest,
		storageHandler.New,
	)

	return &storageHandler.Handler{}, nil
}
