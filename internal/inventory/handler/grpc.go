package inventoryHandler

import (
	"context"

	inventoryService "github.com/hifat/mallow-sale-api/internal/inventory/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/grpc/inventoryProto"
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
	for i, v := range invRes.Items {
		res.Items[i] = &inventoryProto.Inventory{
			Id:   v.ID,
			Name: v.Name,
		}
	}

	return &res, nil
}
