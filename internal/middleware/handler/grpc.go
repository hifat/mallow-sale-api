package middlewareHandler

import (
	"context"

	"github.com/hifat/mallow-sale-api/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Middleware struct {
	cfg *config.GRPC
}

func NewGRPC(cfg *config.GRPC) *Middleware {
	return &Middleware{cfg}
}

func (s *Middleware) AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "messing metadata")
	}

	apiKeys := md.Get("x-api-key")
	if len(apiKeys) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing api key")
	}

	if apiKeys[0] != s.cfg.Key {
		return nil, status.Error(codes.Unauthenticated, "invalid api key")
	}

	return handler(ctx, req)
}
