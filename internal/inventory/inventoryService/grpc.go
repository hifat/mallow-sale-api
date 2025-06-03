package inventoryService

import (
	"context"

	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryRepository"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitRepository"
	"github.com/hifat/mallow-sale-api/pkg/utils"
)

type IInventoryGRPCService interface {
	FindIn(ctx context.Context, filter inventory.FilterReq) (*inventoryProto.InventoryRes, error)
}

type inventoryGRPCService struct {
	logger            core.Logger
	inventoryRepo     inventoryRepository.IInventoryRepository
	usageUnitGRPCRepo usageUnitRepository.IUsageUnitGRPCRepository
}

func NewGRPC(logger core.Logger, inventoryRepo inventoryRepository.IInventoryRepository, usageUnitGRPCRepo usageUnitRepository.IUsageUnitGRPCRepository) IInventoryGRPCService {
	return &inventoryGRPCService{
		logger,
		inventoryRepo,
		usageUnitGRPCRepo,
	}
}

func (s *inventoryGRPCService) FindIn(ctx context.Context, filter inventory.FilterReq) (*inventoryProto.InventoryRes, error) {
	inventories, err := s.inventoryRepo.FindIn(ctx, filter)
	if err != nil {
		return nil, err
	}

	res := make([]*inventoryProto.Inventory, 0, len(inventories))
	for _, v := range inventories {
		item := &inventoryProto.Inventory{
			Id:               v.ID,
			Name:             v.Name,
			PurchasePrice:    v.PurchasePrice,
			PurchaseQuantity: v.PurchaseQuantity,
			YieldPercentage:  v.YieldPercentage,
			Remark:           v.Remark,
			CreatedAt:        utils.MustToTimestamp(v.CreatedAt),
			UpdatedAt:        utils.MustToTimestamp(v.UpdatedAt),
			PurchaseUnit: &inventoryProto.UsageUnitEmbed{
				Code: v.PurchaseUnit.Code,
				Name: v.PurchaseUnit.Name,
			},
		}

		res = append(res, item)
	}

	return &inventoryProto.InventoryRes{
		Items: res,
	}, nil
}
