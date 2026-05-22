package storagesvc

import (
	"context"

	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	fileStatusModule "github.com/hifat/mallow-sale-api/internal/fileStatus"
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type storageService struct {
	cfg         *config.Config
	log         logger.ILogger
	utilsHelper storageModule.IHelper
	grpcRepo    storageModule.IGrpcRepository
	repo        storageModule.IRepository
}

func New(cfg *config.Config, log logger.ILogger, utilsHelper storageModule.IHelper, grpcRepo storageModule.IGrpcRepository, repo storageModule.IRepository) storageModule.IService {
	return &storageService{
		cfg:         cfg,
		log:         log,
		grpcRepo:    grpcRepo,
		utilsHelper: utilsHelper,
		repo:        repo,
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

	res, err := s.grpcRepo.Upload(ctx, &storageproto.UploadRequest{
		File:        req.File,
		Filename:    req.Filename,
		ContentType: req.ContentType,
		Path:        path,
	})
	if err != nil {
		s.log.Error(err)
		return nil, handling.ThrowErr(err)
	}

	createReq := &storageModule.CreateStorageRequest{
		Filename:   req.Filename,
		ObjectKey:  res.ObjectKey,
		StatusCode: fileStatusModule.EnumFileStatusCodeOrphaned,
	}

	createRes, err := s.repo.Create(ctx, createReq)
	if err != nil {
		s.log.Error(err)
		return nil, handling.ThrowErr(err)
	}

	createRes.Url = res.Url

	return &handling.ResponseItem[*storageModule.UploadResponse]{
		Item: createRes,
	}, nil
}
