package rpc

import (
	"fmt"
	"log"
	"net"

	inventoryDi "github.com/hifat/mallow-sale-api/internal/inventory/di"
	inventoryProto "github.com/hifat/mallow-sale-api/internal/inventory/proto"
	middlewareHandler "github.com/hifat/mallow-sale-api/internal/middleware/handler"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterGRPC(cfg *config.Config, db *mongo.Database) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("listening on port :%s", cfg.GRPC.Port)

	m := middlewareHandler.NewGRPC(&cfg.GRPC)

	grpcSrv := grpc.NewServer(
		grpc.UnaryInterceptor(m.AuthInterceptor),
	)

	ivnDi := inventoryDi.Init(cfg, db)

	inventoryProto.RegisterInventoryGrpcServiceServer(grpcSrv, ivnDi.GRPC)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
