package router

import (
	"fmt"
	"log/slog"

	inventoryDI "github.com/hifat/mallow-sale-api/internal/inventory/inventoryDi"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
)

func (r *router) InventoryRouter() {
	handler := inventoryDI.Init(r.cfg, r.db, r.logger, r.validator, r.grpc)

	go func() {
		grpcServer, lis, err := rpc.NewGRPCServer(&r.cfg.Auth, r.cfg.GRPC.InventoryHost)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		inventoryProto.RegisterInventoryGrpcServiceServer(grpcServer, handler.GRPC)

		slog.Info(fmt.Sprintf("Inventory gRPC server listening on: %s", r.cfg.GRPC.InventoryHost))
		grpcServer.Serve(lis)
	}()

	inventory := r.route.Group("/api/inventories")
	inventory.Get("", handler.InventoryRest.Find)
	inventory.Get("/:inventoryID", handler.InventoryRest.FindByID)
	inventory.Post("", handler.InventoryRest.Create)
	inventory.Put("/:inventoryID", handler.InventoryRest.Update)
	inventory.Delete("/:inventoryID", handler.InventoryRest.Delete)
}
