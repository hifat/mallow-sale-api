package shoppingService

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IReceiptService interface {
	Reader(ctx context.Context, req *shoppingModule.ReqReceiptReader) (*handling.ResponseItem[shoppingModule.ResReceiptReader], error)
}

type receiptService struct {
	logger logger.ILogger
}

func NewReceipt(logger logger.ILogger) IReceiptService {
	return &receiptService{
		logger,
	}
}

func (s *receiptService) upload(file *multipart.FileHeader, targetPath string) (*string, error) {
	fileHeader, err := file.Open()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	defer fileHeader.Close()

	mtype, err := mimetype.DetectReader(fileHeader)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	allowedImageType := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
		"image/bmp":  true,
	}

	_, ok := allowedImageType[mtype.String()]
	if !ok {
		return nil, errors.New("invalid mimetype")
	}

	ext := mtype.Extension()
	if ext == "" {
		ext = strings.ToLower(filepath.Ext(file.Filename))
	}

	// Generate unique filename
	filename := filepath.Base(file.Filename)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	newFileName := fmt.Sprintf("%s_%d%s", filename, time.Now().Unix(), ext)

	// Create destination file
	dst, err := os.Create(filepath.Join(targetPath, newFileName))
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	defer dst.Close()

	// Open source file
	src, err := file.Open()
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	defer src.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return nil, nil
}

func (s *receiptService) Reader(ctx context.Context, req *shoppingModule.ReqReceiptReader) (*handling.ResponseItem[shoppingModule.ResReceiptReader], error) {
	var MakeFileSize int64 = 10 * 1024 * 1024 // 10 MB
	if req.Image.Size > MakeFileSize {
		return nil, errors.New("file size must be less than 10 MB")
	}

	_, err := s.upload(req.Image, "upload")
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	return &handling.ResponseItem[shoppingModule.ResReceiptReader]{}, nil
}
