package storageHandler

import (
	"io"

	"github.com/gin-gonic/gin"
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	storageService "github.com/hifat/mallow-sale-api/internal/storage/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	storageService storageService.IService
}

func NewRest(storageService storageService.IService) *Rest {
	return &Rest{storageService: storageService}
}

// @Summary 	Upload File to Google Drive
// @security 	BearerAuth
// @Tags 		storage
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		file formData file true "File to upload"
// @Param 		folderId formData string false "Google Drive folder ID (optional)"
// @Success 	201 {object} handling.ResponseItem[storageModule.UploadResponse]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/storage/upload [post]
func (r *Rest) Upload(c *gin.Context) {
	// Get file from form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		handling.ResponseErr(c, handling.ThrowErr(err))
		return
	}
	defer file.Close()

	// Read file content
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		handling.ResponseErr(c, handling.ThrowErr(err))
		return
	}

	// Get optional folder ID
	folderID := c.PostForm("folderId")

	// Create upload request
	req := &storageModule.UploadRequest{
		File:     fileBytes,
		FileName: header.Filename,
		MimeType: header.Header.Get("Content-Type"),
		FolderID: folderID,
	}

	// Upload file
	res, err := r.storageService.Upload(c.Request.Context(), req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}
