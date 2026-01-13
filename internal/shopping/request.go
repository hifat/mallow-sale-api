package shoppingModule

import (
	"mime/multipart"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type RequestInventoryStatus struct {
	Code EnumCodeInventoryStatusType `json:"code"`
	Name string                      `json:"-"`
}

type RequestInventory struct {
	OrderNo          uint                   `json:"orderNo"`
	InventoryID      string                 `json:"inventoryID"`
	InventoryName    string                 `json:"inventoryName"`
	PurchaseQuantity float64                `json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Entity `json:"purchaseUnit"`
}

type Request struct {
	SupplierID   string             `json:"supplierID"`
	SupplierName string             `json:"-"`
	Inventories  []RequestInventory `json:"inventories"`
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
