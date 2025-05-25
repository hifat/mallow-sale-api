package rpc

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"github.com/hifat/mallow-sale-api/internal/usageUnit/usageUnitProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type grpcAuth struct {
	secretKey string
}

func (g *grpcAuth) unaryAuth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata not found")
	}

	authHandler, ok := md["x-api-key"]
	if !ok {
		return nil, errors.New("x-api-key metadata not found")
	}

	if len(authHandler) == 0 {
		return nil, errors.New("x-api-key is empty value")
	}

	// _, err := jwtauth.ParseToken(g.secretKey, string(authHandler[0]))
	// if err != nil {
	// 	return nil, errors.New("token is invalid")
	// }
	// logger.Info(fmt.Sprintf("claims: %+v", claim))

	return handler(ctx, req)
}

type GrpcClient interface {
	Inventory() inventoryProto.InventoryGrpcServiceClient
	UsageUnit() usageUnitProto.UsageUnitGrpcServiceClient
	CloseAll()
}

type grpcClient struct {
	inventoryConn *grpc.ClientConn
	usageUnitConn *grpc.ClientConn
}

func NewGRPCClient(cfg *config.Config) (GrpcClient, error) {
	// Connection options with a timeout
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Connect to User service
	inventoryConn, err := grpc.NewClient(cfg.GRPC.InventoryHost, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to inventory service: %v", err)
	}

	// Connect to User service
	usageUnitConn, err := grpc.NewClient(cfg.GRPC.UsageUnitHost, opts...)
	if err != nil {
		if err := inventoryConn.Close(); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to connect to usageUnit service: %v", err)
	}

	return &grpcClient{
		inventoryConn,
		usageUnitConn,
	}, nil
}

func (g *grpcClient) Inventory() inventoryProto.InventoryGrpcServiceClient {
	return inventoryProto.NewInventoryGrpcServiceClient(g.inventoryConn)
}

func (g *grpcClient) UsageUnit() usageUnitProto.UsageUnitGrpcServiceClient {
	return usageUnitProto.NewUsageUnitGrpcServiceClient(g.usageUnitConn)
}

func (c *grpcClient) CloseAll() {
	if c.inventoryConn != nil {
		c.inventoryConn.Close()
	}
}

func NewGRPCServer(cfg *config.Auth, host string) (*grpc.Server, net.Listener, error) {
	grpcAuth := &grpcAuth{
		secretKey: cfg.APIKey,
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcAuth.unaryAuth),
	}
	recover()
	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, nil, err
	}

	return grpcServer, lis, nil
}
