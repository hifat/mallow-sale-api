package purchaseSupplierEvidenceModule

import (
	purchaseSupplierEvidenceTypeModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence/type"
)

type CreateEvidenceRequest struct {
	Type      purchaseSupplierEvidenceTypeModule.EnumPurchaseSupplierEvidenceTypeCode `json:"type" binding:"required"`
	FileName  string                                                                  `json:"fileName" binding:"required"`
	ObjectKey string                                                                  `json:"objectKey" binding:"required"`

	PurchaseSupplierID string `json:"-"`
}
