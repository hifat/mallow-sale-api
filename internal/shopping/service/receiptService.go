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

type receiptService struct {
	logger      logger.ILogger
	receiptGRPC shoppingModule.IReceiptGrpcRepository
}

func NewReceipt(logger logger.ILogger, receiptGRPC shoppingModule.IReceiptGrpcRepository) shoppingModule.IReceiptService {
	return &receiptService{
		logger,
		receiptGRPC,
	}
}

func (s *receiptService) upload(file *multipart.FileHeader, targetPath string) (string, error) {
	fileHeader, err := file.Open()
	if err != nil {
		s.logger.Error(err)
		return "", err
	}
	defer fileHeader.Close()

	mtype, err := mimetype.DetectReader(fileHeader)
	if err != nil {
		s.logger.Error(err)
		return "", err
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
		return "", errors.New("invalid mimetype")
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
		return "", err
	}
	defer dst.Close()

	// Open source file
	src, err := file.Open()
	if err != nil {
		s.logger.Error(err)
		return "", err
	}
	defer src.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		s.logger.Error(err)
		return "", err
	}

	return newFileName, nil
}

func (s *receiptService) Reader(ctx context.Context, req *shoppingModule.ReqReceiptReader) (*handling.ResponseItems[shoppingModule.ResReceiptReader], error) {
	var MakeFileSize int64 = 10 * 1024 * 1024 // 10 MB
	if req.Image.Size > MakeFileSize {
		return nil, errors.New("file size must be less than 10 MB")
	}

	imgName, err := s.upload(req.Image, "upload")
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	img, _ := os.ReadFile("upload/" + imgName)
	resRcp, err := s.receiptGRPC.ReadReceipt(ctx, imgName, img)
	if err != nil {
		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	// TODO: Remove file when used

	return &handling.ResponseItems[shoppingModule.ResReceiptReader]{
		Items: resRcp,
	}, nil
}
