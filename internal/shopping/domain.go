package shoppingModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type Request struct {
	Name             string                       `fake:"{name}" json:"name"`
	PurchaseQuantity float64                      `fake:"{float64}" json:"purchaseQuantity"`
	IsComplete       bool                         `fake:"{bool}" json:"isComplete"`
	PurchaseUnit     usageUnitModule.UsageUnitReq `json:"purchaseUnit"`
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
