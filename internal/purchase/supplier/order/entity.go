package purchaseSupplierOrderModule

import (
	purchaseStatusModule "github.com/hifat/mallow-sale-api/internal/purchase/status"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	utilsModule.Base `bson:"inline"`

	ID                 primitive.ObjectID                          `bson:"_id,omitempty" json:"id"`
	PurchaseSupplierID primitive.ObjectID                          `bson:"purchase_supplier_id" json:"purchaseSupplierID"`
	InventoryID        primitive.ObjectID                          `bson:"inventory_id" json:"inventoryID"`
	InventoryName      string                                      `bson:"inventory_name" json:"inventoryName"`
	Quantity           float64                                     `bson:"quantity" json:"quantity"`
	UsageUnitCode      string                                      `bson:"usage_unit_code" json:"usageUnitCode"`
	StatusCode         purchaseStatusModule.EnumPurchaseStatusCode `bson:"status_code" json:"statusCode"`
}
