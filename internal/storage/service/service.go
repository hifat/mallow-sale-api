package storageService

import (
	"context"

	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	storageRepository "github.com/hifat/mallow-sale-api/internal/storage/repository"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Upload(ctx context.Context, req *storageModule.UploadRequest) (*handling.ResponseItem[*storageModule.UploadResponse], error)
}

type service struct {
	logger      logger.ILogger
	storageRepo storageRepository.IRepository
}

func New(
	logger logger.ILogger,
	storageRepo storageRepository.IRepository,
) IService {
	return &service{
		logger:      logger,
		storageRepo: storageRepo,
	}
}

func (s *service) Upload(ctx context.Context, req *storageModule.UploadRequest) (*handling.ResponseItem[*storageModule.UploadResponse], error) {
	// Upload file to Google Drive
	uploadedFile, err := s.storageRepo.Upload(ctx, req)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*storageModule.UploadResponse]{
		Item: uploadedFile,
	}, nil
}
