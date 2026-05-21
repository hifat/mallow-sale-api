//go:build wireinject
// +build wireinject

package storageDi

import (
	"github.com/google/wire"
	storageHandler "github.com/hifat/mallow-sale-api/internal/storage/handler"
	storageHelper "github.com/hifat/mallow-sale-api/internal/storage/helper"
	storageRepository "github.com/hifat/mallow-sale-api/internal/storage/repository"
	storageService "github.com/hifat/mallow-sale-api/internal/storage/service"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"google.golang.org/grpc"
)

//
//go:generate wire ./wire.go
func Init(cfg *config.Config, grpcConn *grpc.ClientConn) *storageHandler.Handler {
	wire.Build(
		logger.New,

		// Repository
		storageRepository.NewGrpcRepository,

		// Helper
		storageHelper.New,

		// Service
		storageService.New,

		// Handler
		storageHandler.NewRest,
		storageHandler.New,
	)

	return &storageHandler.Handler{}
}
