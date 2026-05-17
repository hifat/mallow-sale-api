package purchaseSupplierModule

import (
	paymentTypeModule "github.com/hifat/mallow-sale-api/internal/paymentType"
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	ID                 primitive.ObjectID                          `bson:"_id,omitempty" json:"id"`
	PurchaseID         primitive.ObjectID                          `bson:"purchase_id" json:"purchaseID"`
	SupplierID         primitive.ObjectID                          `bson:"supplier_id" json:"supplierID"`
	SupplierName       string                                      `bson:"supplier_name" json:"supplierName"`
	PurchaseStatusCode purchaseStatusModule.EnumPurchaseStatusCode `bson:"purchase_status_code" json:"purchaseStatusCode"`
	PaymentTypeCode    paymentTypeModule.EnumPaymentTypeCode       `bson:"payment_type_code" json:"paymentTypeCode"`
}
