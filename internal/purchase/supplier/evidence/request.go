package purchaseSupplierEvidenceModule

import (
	"time"

	fileStatusModule "github.com/hifat/mallow-sale-api/internal/fileStatus"
	purchaseSupplierEvidenceTypeModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence/type"
)

type CreateEvidenceRequest struct {
	Type           purchaseSupplierEvidenceTypeModule.EnumPurchaseSupplierEvidenceTypeCode `json:"type" binding:"required"`
	FileName       string                                                                  `json:"file_name" binding:"required"`
	FileRename     string                                                                  `json:"file_rename" binding:"required"`
	Path           string                                                                  `json:"path" binding:"required"`
	FileStatusCode fileStatusModule.EnumFileStatusCode                                     `json:"file_status" binding:"required"`
	UploadedAt     time.Time                                                               `json:"uploaded_at" binding:"required"`

	PurchaseSupplierID string `json:"-"`
}
