package inventoryHandler

import (
	"context"

	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryService"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	filter := inventory.FilterReq{}
	if err := copier.Copy(&filter, req); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	inventories, err := h.inventorySrv.FindIn(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
	}, nil
}
