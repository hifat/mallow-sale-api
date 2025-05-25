package inventoryRepository

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/entity"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
	"github.com/hifat/mallow-sale-api/pkg/utils/repoUtils"
	"google.golang.org/grpc/metadata"
)

type IInventoryGRPCRepository interface {
	FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.Inventory, error)
}

type grpcRepository struct {
	grpcClient rpc.GrpcClient
}

func NewGRPC(grpcClient rpc.GrpcClient) IInventoryGRPCRepository {
	return &grpcRepository{grpcClient}
}

func (r *grpcRepository) FindIn(ctx context.Context, filter inventory.FilterReq) ([]inventory.Inventory, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-api-key", "sdfsdf"))
	result, err := r.grpcClient.Inventory().FindIn(ctx, &inventoryProto.InFilter{
		Ids:   filter.IDs,
		Codes: filter.Codes,
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
		_inventory.PurchaseUnit = &usageUnit.UsageUnitEmbed{}
		_inventory.PurchaseUnit.SetAttr(purchaseUnit.Code, purchaseUnit.Name)

		inventories = append(inventories, _inventory)
	}

	return inventories, nil
}
