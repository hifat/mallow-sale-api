package inventoryHandler

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryService"
)

type inventoryGRPC struct {
	inventoryProto.UnimplementedInventoryGrpcServiceServer
	inventorySrv inventoryService.IInventoryService
}

func NewGRPC(inventorySrv inventoryService.IInventoryService) *inventoryGRPC {
	return &inventoryGRPC{
		inventorySrv: inventorySrv,
	}
}

// TODO: Reflect and make safe this func
func (h *inventoryGRPC) FindIn(ctx context.Context, req *inventoryProto.InFilter) (*inventoryProto.InventoryRes, error) {
	inventories, err := h.inventorySrv.FindInID(ctx, req.Ids)

	res := make([]*inventoryProto.Inventory, 0, len(inventories))
	for _, v := range inventories {
		item := &inventoryProto.Inventory{
			Id:               v.ID,
			Name:             v.Name,
			PurchasePrice:    v.PurchasePrice,
			PurchaseQuantity: v.PurchaseQuantity,
			YieldPercentage:  v.YieldPercentage,
			Remark:           v.Remark,
			// CreatedAt:        v.CreatedAt,
			// UpdatedAt:        v.UpdatedAt,
			PurchaseUnit: &inventoryProto.UsageUnitEmbed{
				Code: v.PurchaseUnit.Code,
				Name: v.PurchaseUnit.Name,
			},
		}

		res = append(res, item)
	}

	return &inventoryProto.InventoryRes{
		Items: res,
	}, err
}
