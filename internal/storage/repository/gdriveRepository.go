package storageRepository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

//go:generate mockgen -source=./gdriveRepository.go -destination=./mock/repository.go -package=mockStorageRepository
type IRepository interface {
	Upload(ctx context.Context, req *storageModule.UploadRequest) (*storageModule.UploadResponse, error)
}

type gdriveRepository struct {
	driveService *drive.Service
}

func NewGDrive(cfg *config.Config) (IRepository, error) {
	ctx := context.Background()

	var err error

	// Try to use the value as JSON string first
	clientSecret := strings.TrimSpace(cfg.GoogleDrive.ServiceAccount)
	if clientSecret == "" {
		return nil, errors.New("GDRIVE_SERVICE_ACCOUNT is not configured (must provide Client Secret JSON)")
	}

	driveService, err := drive.NewService(ctx, option.WithAuthCredentialsJSON(option.AuthorizedUser, []byte(clientSecret)))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Drive client: %w", err)
	}

	return &gdriveRepository{
		driveService: driveService,
	}, nil
}

func (r *gdriveRepository) Upload(ctx context.Context, req *storageModule.UploadRequest) (*storageModule.UploadResponse, error) {
	// Create file metadata
	file := &drive.File{
		Name:     req.FileName,
		MimeType: req.MimeType,
		Parents:  []string{""},
	}

	// Create reader from file bytes
	reader := io.NopCloser(bytes.NewReader(req.File))

	// Upload file to Google Drive
	uploadedFile, err := r.driveService.Files.Create(file).
		Context(ctx).
		Media(reader).
		Fields("id, name, webViewLink").
		Do()
	if err != nil {
		return nil, fmt.Errorf("gdrive upload file error: %s", err)
	}

	now := time.Now()
	return &storageModule.UploadResponse{
		FileID:      uploadedFile.Id,
		FileName:    uploadedFile.Name,
		WebViewLink: uploadedFile.WebViewLink,
		UploadedAt:  &now,
	}, nil
}
