package storageModule

import (
	fileStatusModule "github.com/hifat/mallow-sale-api/internal/fileStatus"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	FileName   string                              `bson:"file_name" json:"fileName"`
	ObjectKey  string                              `bson:"object_key" json:"objectKey"`
	StatusCode fileStatusModule.EnumFileStatusCode `bson:"status_code" json:"statusCode"`
}
