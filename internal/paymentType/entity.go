package paymentTypeModule

import (
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	ID   primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Code EnumPaymentTypeCode `bson:"code" json:"code"`
	Name string              `bson:"name" json:"name"`
}
