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
	inventoryGRPCSrv inventoryService.IInventoryGRPCService
}

func NewGRPC(inventoryGRPCSrv inventoryService.IInventoryGRPCService) *inventoryGRPC {
	return &inventoryGRPC{
		inventoryGRPCSrv: inventoryGRPCSrv,
	}
}

func (h *inventoryGRPC) FindIn(ctx context.Context, req *inventoryProto.InFilter) (*inventoryProto.InventoryRes, error) {
	filter := inventory.FilterReq{}
	if err := copier.Copy(&filter, req); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return h.inventoryGRPCSrv.FindIn(ctx, filter)
}
