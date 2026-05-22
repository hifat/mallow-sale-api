package storageModule

import fileStatusModule "github.com/hifat/mallow-sale-api/internal/fileStatus"

type UploadRequest struct {
	File        []byte `json:"file"`
	ServiceCode string `json:"serviceCode"`

	ContentType string `json:"-"`
	Filename    string `json:"-"`
}

type CreateStorageRequest struct {
	Filename   string
	ObjectKey  string
	StatusCode fileStatusModule.EnumFileStatusCode
}
