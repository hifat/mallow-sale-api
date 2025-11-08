package userModule

import utilsModule "github.com/hifat/mallow-sale-api/internal/utils"

type Entity struct {
	utilsModule.Base

	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"-"`
}
