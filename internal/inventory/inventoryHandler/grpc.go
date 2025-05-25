package inventoryHandler

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryService"
	"github.com/hifat/mallow-sale-api/pkg/utils"
	"github.com/jinzhu/copier"
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

func (h *inventoryGRPC) FindIn(ctx context.Context, req *inventoryProto.InFilter) (*inventoryProto.InventoryRes, error) {
	filter := inventory.FilterReq{}
	if err := copier.Copy(&filter, req); err != nil {
		return nil, err
	}

	inventories, err := h.inventorySrv.FindIn(ctx, filter)
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
