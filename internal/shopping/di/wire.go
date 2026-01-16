//go:build wireinject
// +build wireinject

package shoppingDi

import (
	"github.com/google/wire"
	inventoryHelper "github.com/hifat/mallow-sale-api/internal/inventory/helper"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	shoppingHandler "github.com/hifat/mallow-sale-api/internal/shopping/handler"
	shoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository"
	shoppingService "github.com/hifat/mallow-sale-api/internal/shopping/service"
	supplierHelper "github.com/hifat/mallow-sale-api/internal/supplier/helper"
	supplierRepository "github.com/hifat/mallow-sale-api/internal/supplier/repository"
	usageUnitHelper "github.com/hifat/mallow-sale-api/internal/usageUnit/helper"
	usageUnitRepository "github.com/hifat/mallow-sale-api/internal/usageUnit/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

//
//go:generate wire ./wire.go
func Init(cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) *shoppingHandler.Handler {
	wire.Build(
		// Repository
		shoppingRepository.NewMongo,
		shoppingRepository.NewInventoryMongo,
		shoppingRepository.NewReceiptGRPC,
		usageUnitRepository.NewMongo,
		supplierRepository.NewMongo,
		inventoryRepository.NewMongo,

		// Service
		logger.New,
		shoppingService.New,
		shoppingService.NewInventory,
		shoppingService.NewReceipt,

		// Handler
		shoppingHandler.NewReceiptRest,
		shoppingHandler.NewInventoryRest,
		shoppingHandler.NewRest,
		shoppingHandler.New,

		// Helper
		usageUnitHelper.New,
		supplierHelper.New,
		inventoryHelper.New,
	)

	return &shoppingHandler.Handler{}
}
