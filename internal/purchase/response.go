package purchaseModule

import (
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
	"time"
)

type Response struct {
	ID                 string                            `json:"id"`
	PurchaseStatusCode purchaseStatusModule.EnumPurchaseStatusCode `json:"purchaseStatusCode"`
	Suppliers          []purchaseSupplierModule.Response `json:"suppliers"`
	CreatedAt          time.Time                         `json:"createdAt"`
	UpdatedAt          time.Time                         `json:"updatedAt"`
}
