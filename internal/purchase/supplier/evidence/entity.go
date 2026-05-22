package purchaseSupplierEvidenceModule

import (
	"time"

	purchaseSupplierEvidenceTypeModule "github.com/hifat/mallow-sale-api/internal/purchase/supplier/evidence/type"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	ID                               primitive.ObjectID                                                      `bson:"_id,omitempty" json:"id"`
	PurchaseSupplierID               primitive.ObjectID                                                      `bson:"purchase_supplier_id" json:"purchaseSupplierID"`
	PurchaseSupplierEvidenceTypeCode purchaseSupplierEvidenceTypeModule.EnumPurchaseSupplierEvidenceTypeCode `bson:"purchase_supplier_evidence_type_code" json:"purchaseSupplierEvidenceTypeCode"`
	FileName                         string                                                                  `bson:"file_name" json:"fileName"`
	ObjectKey                        string                                                                  `bson:"object_key" json:"objectKey"`
	UploadedAt                       time.Time                                                               `bson:"uploaded_at" json:"uploadedAt"`
}
