package shoppingModule

import (
	"mime/multipart"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type RequestInventoryStatus struct {
	Code EnumCodeInventoryStatusType `json:"code"`

	Name string `json:"-"`
}

type RequestInventory struct {
	OrderNo          uint                         `validate:"required" json:"orderNo"`
	InventoryID      string                       `validate:"required" json:"inventoryID"`
	PurchaseQuantity float64                      `validate:"required" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.ReqUsageUnit `json:"purchaseUnit"`

	InventoryName string          `json:"-"`
	Status        InventoryStatus `json:"-"`
}

type Request struct {
	SupplierID  string             `validate:"required" json:"supplierID"`
	Inventories []RequestInventory `validate:"dive" json:"inventories"`

	SupplierName string `json:"-"`
}

func (p *Request) GetPurchaseUnitCodes() []string {
	appended := map[string]bool{}
	codes := make([]string, 0, len(p.Inventories))
	for _, v := range p.Inventories {
		if ok := appended[v.PurchaseUnit.Code]; !ok {
			appended[v.PurchaseUnit.Code] = true
			codes = append(codes, v.PurchaseUnit.Code)
		}
	}

	return codes
}

func (p *Request) GetInventoryIDs() []string {
	ids := make([]string, 0, len(p.Inventories))
	for _, v := range p.Inventories {
		ids = append(ids, v.InventoryID)
	}

	return ids
}

type ReqReOrder struct {
	ID      string `fake:"{uuid}" json:"id"`
	OrderNo uint   `fake:"{uintrange:0,100}" json:"orderNo"`
}

type ReqUpdateIsComplete struct {
	IsComplete bool `json:"isComplete"`
}

/* --------------------------------- Receipt -------------------------------- */

type ReqReceiptReader struct {
	Image *multipart.FileHeader `json:"image"`
}
