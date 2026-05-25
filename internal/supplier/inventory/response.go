package supplierInventoryModule

import (
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
)

type Response struct {
	supplierModule.Prototype
	Inventories []inventoryModule.Prototype `json:"inventories"`
}

type GroupBySupplierResponse struct {
	supplierModule.Prototype `json:"supplier"`
	Inventories              []inventoryModule.Response `json:"inventories"`
}
