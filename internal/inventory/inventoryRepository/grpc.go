package inventoryRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"github.com/hifat/mallow-sale-api/pkg/utils/repoUtils"
)

type IInventoryGRPCRepository interface {
	FindInID(ctx context.Context, ids []string) ([]inventory.Inventory, error)
}

type grpcRepository struct {
	grpcClient rpc.GrpcClient
}

func NewGRPC(grpcClient rpc.GrpcClient) IInventoryGRPCRepository {
	return &grpcRepository{grpcClient}
}

func (r *grpcRepository) FindInID(ctx context.Context, ids []string) ([]inventory.Inventory, error) {
	result, err := r.grpcClient.Inventory().FindIn(ctx, &inventoryProto.InFilter{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	inventories := make([]inventory.Inventory, 0, len(result.Items))
	for _, v := range result.Items {
		_inventory := inventory.Inventory{
			Base: entity.Base{
				ID:        v.Id,
				CreatedAt: repoUtils.TimestampToTimePtr(v.CreatedAt),
				UpdatedAt: repoUtils.TimestampToTimePtr(v.UpdatedAt),
			},
			Name:             v.Name,
			PurchasePrice:    v.PurchasePrice,
			PurchaseQuantity: v.PurchaseQuantity,
			YieldPercentage:  v.YieldPercentage,
			Remark:           v.Remark,
		}

		purchaseUnit := v.PurchaseUnit
		_inventory.PurchaseUnit.SetAttr(purchaseUnit.Code, purchaseUnit.Name)

		inventories = append(inventories, _inventory)
	}

	return inventories, nil
}
