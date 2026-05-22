package purchaseSupplierEvidenceModule

import (
	"time"

	fileStatusModule "github.com/hifat/mallow-sale-api/internal/fileStatus"
	purchaseSupplierEvidenceTypeModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence/type"
)

type Response struct {
	ID                               string                                                                  `json:"id"`
	PurchaseSupplierID               string                                                                  `json:"purchaseSupplierID"`
	PurchaseSupplierEvidenceTypeCode purchaseSupplierEvidenceTypeModule.EnumPurchaseSupplierEvidenceTypeCode `json:"purchaseSupplierEvidenceTypeCode"`
	FileName                         string                                                                  `json:"fileName"`
	ObjectKey                        string                                                                  `json:"objectKey"`
	FileStatusCode                   fileStatusModule.EnumFileStatusCode                                     `json:"fileStatusCode"`
	UploadedAt                       time.Time                                                               `json:"uploadedAt"`
	CreatedAt                        time.Time                                                               `json:"createdAt"`
	UpdatedAt                        time.Time                                                               `json:"updatedAt"`
}
