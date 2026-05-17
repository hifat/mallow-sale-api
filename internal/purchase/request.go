package purchaseModule

import (
	purchaseSupplierModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier"
)

type CreatePurchaseRequest struct {
	Suppliers []purchaseSupplierModule.CreateSupplierRequest `json:"suppliers" binding:"required,dive,required"`

	ID string `json:"-"`
}
