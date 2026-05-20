package purchaseSupplierModule

import (
	paymentTypeModule "github.com/hifat/mallow-sale-api/internal/paymentType"
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	purchaseSupplierOrderModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/order"
)

type CreateSupplierRequest struct {
	SupplierID   string                                           `json:"supplierId" binding:"required"`
	SupplierName string                                           `json:"supplierName" binding:"required"`
	StatusCode   purchaseStatusModule.EnumPurchaseStatusCode      `json:"statusCode" binding:"required"`
	PaymentType  paymentTypeModule.EnumPaymentTypeCode            `json:"paymentType" binding:"required"`
	Orders       []purchaseSupplierOrderModule.CreateOrderRequest `json:"orders" binding:"required,dive,required"`

	PurchaseSupplierID string `json:"-"`
}
