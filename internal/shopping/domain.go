package shoppingModule

import (
	"mime/multipart"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type Request struct {
	Name             string                       `fake:"{name}" json:"name"`
	PurchaseQuantity float64                      `fake:"{float64}" json:"purchaseQuantity"`
	IsComplete       bool                         `fake:"{bool}" json:"isComplete"`
	PurchaseUnit     usageUnitModule.UsageUnitReq `json:"purchaseUnit"`
}

type ReqReOrder struct {
	ID      string `fake:"{uuid}" json:"id"`
	OrderNo uint   `fake:"{uintrange:0,100}" json:"orderNo"`
}

type Response struct {
	ID               string                    `json:"id"`
	Name             string                    `fake:"{name}" json:"name"`
	IsComplete       bool                      `fake:"{bool}" json:"isComplete"`
	PurchaseQuantity float64                   `fake:"{float64}" json:"purchaseQuantity"`
	PurchaseUnit     usageUnitModule.Prototype `json:"purchaseUnit"`
}

type ReqUpdateIsComplete struct {
	IsComplete bool `json:"isComplete"`
}

type ReqReceiptReader struct {
	Image *multipart.FileHeader `json:"image"`
}

type ResReceiptReader struct {
	InventoryID      string  `json:"inventoryID"` // Make vector db solution for search
	Name             string  `json:"name"`
	NameEdited       string  `json:"nameEdited"`
	PurchasePrice    float64 `json:"purchasePrice"`
	PurchaseQuantity float64 `json:"purchaseQuantity"`
	Remark           string  `json:"remark"`
}
