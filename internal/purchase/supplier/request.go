package purchaseSupplierModule

import (
	paymentTypeModule "github.com/hifat/mallow-sale-api/internal/paymentType"
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	purchaseSupplierEvidenceModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
)

type CreateSupplierRequest struct {
	SupplierID   string                                                 `json:"supplier_id" binding:"required"`
	SupplierName string                                                 `json:"supplier_name" binding:"required"`
	Status       purchaseStatusModule.EnumPurchaseStatusCode            `json:"status" binding:"required"`
	PaymentType  paymentTypeModule.EnumPaymentTypeCode                  `json:"payment_type" binding:"required"`
	Orders       []purchaseSupplierOrderModule.CreateOrderRequest       `json:"orders" binding:"required,dive,required"`
	Evidences    []purchaseSupplierEvidenceModule.CreateEvidenceRequest `json:"evidences" binding:"required,dive,required"`

	PurchaseSupplierID string `json:"-"`
}
