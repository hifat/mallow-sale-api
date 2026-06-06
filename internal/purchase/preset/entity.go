package purchasePresetModule

import (
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryEntity struct {
	ID               string `json:"id"`
	No               uint   `bson:"no" json:"no"`
	Name             string `bson:"name" json:"name"`
	PurchaseUnitCode string `bson:"purchase_unit_code" json:"purchaseUnitCode"`
}

type Entity struct {
	utilsModule.Base `bson:"inline"`

	SupplierID   primitive.ObjectID `bson:"supplier_id" json:"supplierID"`
	SupplierName string             `bson:"supplier_name" json:"supplierName"`
	Inventories  []InventoryEntity  `bson:"inventories" json:"inventories"`
}
