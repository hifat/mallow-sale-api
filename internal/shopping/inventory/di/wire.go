//go:build wireinject
// +build wireinject

package shoppingDi

import (
	"github.com/google/wire"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	shoppingInventoryHandler "github.com/hifat/mallow-sale-api/internal/shopping/handler"
	shoppingInventoryRepository "github.com/hifat/mallow-sale-api/internal/shopping/inventory/repository"
	shoppingInventoryService "github.com/hifat/mallow-sale-api/internal/shopping/inventory/service"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Init(cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) *shoppingInventoryHandler.Handler {
	wire.Build(
		// Repository
		shoppingInventoryRepository.NewMongo,
		supplierRepository.NewMongo,
		inventoryRepository.NewMongo,

		// Service
		logger.New,
		shoppingInventoryService.New,

		// Handler
		shoppingInventoryHandler.NewRest,
		shoppingInventoryHandler.New,
	)

	return &shoppingInventoryHandler.Handler{}
}
