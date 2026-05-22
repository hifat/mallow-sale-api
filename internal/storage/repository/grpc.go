package storagerepo

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/config"
)

type grpcRepository struct {
	cfg      *config.Config
	grpcConn *grpc.ClientConn
}

func NewGrpc(cfg *config.Config, grpcConn *grpc.ClientConn) storageModule.IGrpcRepository {
	return &grpcRepository{
		cfg:      cfg,
		grpcConn: grpcConn,
	}
}

func (r *grpcRepository) getClient() storageproto.StorageClient {
	return storageproto.NewStorageClient(r.grpcConn)
}

func (r *grpcRepository) addAuthMetadata(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "x-api-key", r.cfg.GRPCStorage.APIKey)
}

func (r *grpcRepository) Upload(ctx context.Context, req *storageproto.UploadRequest) (*storageproto.UploadResponse, error) {
	ctx = r.addAuthMetadata(ctx)
	return r.getClient().Upload(ctx, req)
}

func (r *grpcRepository) GetPresignedURL(ctx context.Context, req *storageproto.GetPresignedURLRequest) (*storageproto.GetPresignedURLResponse, error) {
	ctx = r.addAuthMetadata(ctx)
	return r.getClient().GetPresignedURL(ctx, req)
}

func (r *grpcRepository) Delete(ctx context.Context, req *storageproto.DeleteRequest) (*storageproto.DeleteResponse, error) {
	ctx = r.addAuthMetadata(ctx)
	return r.getClient().Delete(ctx, req)
}
