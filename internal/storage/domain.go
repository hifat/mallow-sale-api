package storageModule

import (
	"context"

	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type IRepository interface {
	Create(ctx context.Context, req *CreateStorageRequest) (*UploadResponse, error)
}

type IGrpcRepository interface {
	Upload(ctx context.Context, req *storageproto.UploadRequest) (*storageproto.UploadResponse, error)
	GetPresignedURL(ctx context.Context, req *storageproto.GetPresignedURLRequest) (*storageproto.GetPresignedURLResponse, error)
	Delete(ctx context.Context, req *storageproto.DeleteRequest) (*storageproto.DeleteResponse, error)
}

type IService interface {
	Upload(ctx context.Context, req *UploadRequest) (*handling.ResponseItem[*UploadResponse], error)
}

type IHelper interface {
	GetDirName(serviceCode string) (string, error)
}
