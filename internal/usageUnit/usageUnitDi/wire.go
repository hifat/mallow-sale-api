//go:build wireinject
// +build wireinject

package usageUnitDi

import (
	"github.com/google/wire"
	"github.com/hifat/goroger-core/helper"
	"github.com/hifat/goroger-core/logger"
	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitHandler"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitService"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var RepoSet = wire.NewSet(
	usageUnitRepository.NewMongo,
)

var ServiceSet = wire.NewSet(
	helper.New,
	logger.New,
	usageUnitService.New,
)

var HandlerSet = wire.NewSet(
	usageUnitHandler.New,
	usageUnitHandler.NewGRPC,
)

func Init(cfg *config.Config, db *mongo.Database, log *zap.Logger, grpc rpc.GrpcClient) usageUnitHandler.Handler {
	wire.Build(
		RepoSet,
		ServiceSet,
		HandlerSet,
	)

	return usageUnitHandler.Handler{}
}
