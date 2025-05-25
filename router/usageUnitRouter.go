package router

import (
	"fmt"
	"log/slog"

	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitDi"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitProto"
	"github.com/hifat/mallow-sale-api/pkg/rpc"
)

func (r *router) UsageUnitRouter() {
	handler := usageUnitDi.Init(r.cfg, r.db, r.logger, r.grpc)

	go func() {
		grpcServer, lis, err := rpc.NewGRPCServer(&r.cfg.Auth, r.cfg.GRPC.UsageUnitHost)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		usageUnitProto.RegisterUsageUnitGrpcServiceServer(grpcServer, handler.GRPC)
		slog.Info(fmt.Sprintf("UsageUnit gRPC server listening on: %s", r.cfg.GRPC.UsageUnitHost))
		grpcServer.Serve(lis)
	}()
}
