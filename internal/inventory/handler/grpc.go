package inventoryHandler

import (
	"context"

	inventoryProto "github.com/hifat/mallow-sale-api/internal/inventory/proto"
	inventoryService "github.com/hifat/mallow-sale-api/internal/inventory/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type GRPC struct {
	inventoryProto.UnimplementedInventoryGrpcServiceServer
	inventorySvc inventoryService.IService
}

func NewGrpc(inventorySvc inventoryService.IService) *GRPC {
	return &GRPC{inventorySvc: inventorySvc}
}

func (g *GRPC) Find(ctx context.Context, req *inventoryProto.Query) (*inventoryProto.InventoryResponse, error) {
	invRes, err := g.inventorySvc.Find(ctx, &utilsModule.QueryReq{})
	if err != nil {
		return nil, err
	}

	res := inventoryProto.InventoryResponse{}
	res.Items = make([]*inventoryProto.Inventory, len(invRes.Items))
	for i, v := range res.Items {
		res.Items[i] = &inventoryProto.Inventory{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	return &res, nil
}
