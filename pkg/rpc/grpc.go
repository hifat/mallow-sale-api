package rpc

import (
	"context"
	"errors"
	"net"

	"github.com/hifat/mallow-sale-api/config"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type GrpcClientFactoryHandler interface {
	Inventory() inventoryProto.InventoryGrpcServiceClient
}

type grpcClientFactory struct {
	client *grpc.ClientConn
}

func (g *grpcClientFactory) Inventory() inventoryProto.InventoryGrpcServiceClient {
	return inventoryProto.NewInventoryGrpcServiceClient(g.client)
}

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

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.NewClient(host, opts...)
	if err != nil {
		return nil, err
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Auth, host string) (*grpc.Server, net.Listener, error) {
	opts := make([]grpc.ServerOption, 0)

	grpcAuth := &grpcAuth{
		secretKey: cfg.APIKey,
	}

	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuth))

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		return nil, nil, err
	}

	return grpcServer, lis, nil
}
