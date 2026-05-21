package storagesvc

import (
	"context"

	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type storageService struct {
	cfg         *config.Config
	log         logger.ILogger
	grpcRepo    storageModule.IGrpcRepository
	utilsHelper storageModule.IHelper
}

func New(cfg *config.Config, log logger.ILogger, grpcRepo storageModule.IGrpcRepository, utilsHelper storageModule.IHelper) storageModule.IService {
	return &storageService{
		cfg:         cfg,
		log:         log,
		grpcRepo:    grpcRepo,
		utilsHelper: utilsHelper,
	}
}

func (s *storageService) Upload(ctx context.Context, req *storageModule.UploadRequest) (*handling.ResponseItem[*storageModule.UploadResponse], error) {
	path, err := s.utilsHelper.GetDirName(req.ServiceCode)
	if err != nil {
		s.log.Error(err)
		return nil, handling.ThrowErr(err)
	}

	// File does not more than 2MB
	if len(req.File) > 2*1024*1024 {
		return nil, handling.ThrowErrByCode(define.CodeFileTooLarge)
	}

	resp, err := s.grpcRepo.Upload(ctx, &storageproto.UploadRequest{
		File:        req.File,
		Filename:    req.Filename,
		ContentType: req.ContentType,
		Path:        path,
	})
	if err != nil {
		s.log.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[*storageModule.UploadResponse]{
		Item: &storageModule.UploadResponse{
			Url: resp.Url,
		},
	}, nil
}
