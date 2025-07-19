package supplierModule

import (
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`
	Name             string `bson:"name" json:"name"`
	ImgUrl           string `bson:"img_url" json:"imgUrl"`
}
