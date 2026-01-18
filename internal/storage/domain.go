package storageModule

import "time"

type UploadRequest struct {
	File     []byte `json:"-"`
	FileName string `json:"fileName" validate:"required"`
	MimeType string `json:"mimeType" validate:"required"`
	FolderID string `json:"folderId"`
}

type UploadResponse struct {
	FileID      string     `json:"fileId"`
	FileName    string     `json:"fileName"`
	WebViewLink string     `json:"webViewLink"`
	UploadedAt  *time.Time `json:"uploadedAt"`
}
