package shoppingModule

import (
	"mime/multipart"

	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
	"github.com/hifat/mallow-sale-api/pkg/define"
)

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
	Status       Status `json:"-"`
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

type ReqUpdateStatus struct {
	StatusCode EnumCodeShoppingStatusType `validate:"required,oneof=PENDING IN_PROGRESS SUCCESS CANCEL" json:"statusCode"`
}

func (r *ReqUpdateStatus) ValidateStatusCode() error {
	switch r.StatusCode {
	case EnumCodeShoppingPending,
		EnumCodeShoppingInProgress,
		EnumCodeShoppingSuccess,
		EnumCodeShoppingCancel:
		return nil
	default:
		return define.ErrInvalidShoppingStatus
	}
}

/* --------------------------- Shopping Inventory --------------------------- */

type RequestShoppingInventory struct {
	InventoryID   string `validate:"required" json:"inventoryID"`
	InventoryName string `validate:"required" json:"inventoryName"`
	SupplierID    string `validate:"required" json:"supplierID"`
	SupplierName  string `validate:"required" json:"supplierName"`
}

/* --------------------------------- Receipt -------------------------------- */

type ReqReceiptReader struct {
	Image *multipart.FileHeader `json:"image"`
}
