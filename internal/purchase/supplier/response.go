package purchaseSupplierModule

import (
	paymentTypeModule "github.com/hifat/mallow-sale-api/internal/paymentType"
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
	"time"
)

type Response struct {
	ID                 string                                       `json:"id"`
	PurchaseID         string                                       `json:"purchaseID"`
	SupplierID         string                                       `json:"supplierID"`
	SupplierName       string                                       `json:"supplierName"`
	PurchaseStatusCode purchaseStatusModule.EnumPurchaseStatusCode  `json:"purchaseStatusCode"`
	PaymentTypeCode    paymentTypeModule.EnumPaymentTypeCode        `json:"paymentTypeCode"`
	Orders             []purchaseSupplierOrderModule.Response       `json:"orders"`
	Evidences          []purchaseSupplierEvidenceModule.Response    `json:"evidences"`
	CreatedAt          time.Time                                    `json:"createdAt"`
	UpdatedAt          time.Time                                    `json:"updatedAt"`
}
