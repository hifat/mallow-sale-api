//go:build wireinject
// +build wireinject

package purchaseDi

import (
	"github.com/google/wire"
	purchaseHandler "github.com/hifat/mallow-sale-api/internal/purchase/handler"
	purchaseRepository "github.com/hifat/mallow-sale-api/internal/purchase/repository"
	purchaseService "github.com/hifat/mallow-sale-api/internal/purchase/service"
	purchaseSupplierEvidenceRepository "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence/repository"
	purchaseSupplierOrderRepository "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order/repository"
	purchaseSupplierRepository "github.com/hifat/mallow-sale-api/internal/purchase/supplier/repository"
	utilsRepository "github.com/hifat/mallow-sale-api/internal/utils/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(cfg *config.Config, db *mongo.Database) *purchaseHandler.Handler {
	wire.Build(
		// Repository
		purchaseRepository.NewMongo,
		purchaseSupplierRepository.NewMongo,
		purchaseSupplierOrderRepository.NewMongo,
		purchaseSupplierEvidenceRepository.NewMongo,
		utilsRepository.NewMongo,

		// Service
		purchaseService.New,

		// Handler
		purchaseHandler.NewRest,
		purchaseHandler.New,
	)

	return &purchaseHandler.Handler{}
}
