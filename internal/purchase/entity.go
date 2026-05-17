package purchaseModule

import (
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	PurchaseStatusCode purchaseStatusModule.EnumPurchaseStatusCode `bson:"purchase_status_code" json:"purchaseStatusCode"`
}
