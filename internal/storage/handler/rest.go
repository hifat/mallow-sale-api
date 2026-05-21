package storageHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	storageService storageModule.IService
}

func NewRest(storageService storageModule.IService) *Rest {
	return &Rest{storageService: storageService}
}

// @Summary 	Upload File
// @security 	BearerAuth
// @Tags 		storage
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		file formData file true "File to upload"
// @Param 		serviceCode formData string true "Service code"
// @Success 	200 {object} handling.ResponseItem[storageModule.UploadResponse]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/uploads [post]
func (r *Rest) Upload(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		handling.ResponseErr(c, fmt.Errorf("failed to get file: %w", err))
		return
	}

	file, err := header.Open()
	if err != nil {
		handling.ResponseErr(c, fmt.Errorf("failed to open file: %w", err))
		return
	}
	defer file.Close()

	fileBytes := make([]byte, header.Size)
	if _, err := file.Read(fileBytes); err != nil {
		handling.ResponseErr(c, fmt.Errorf("failed to read file: %w", err))
		return
	}

	serviceCode := c.PostForm("serviceCode")

	req := &storageModule.UploadRequest{
		File:        fileBytes,
		Filename:    header.Filename,
		ContentType: header.Header.Get("Content-Type"),
		ServiceCode: serviceCode,
	}

	res, err := r.storageService.Upload(c.Request.Context(), req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}
