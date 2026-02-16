package inventoryModule

import usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"

type Request struct {
	Name            string                       `fake:"{name}" validate:"required" json:"name"`
	PurchaseUnit    usageUnitModule.UsageUnitReq `json:"purchaseUnit"`
	YieldPercentage float32                      `fake:"{float32}" validate:"required" json:"yieldPercentage"`
	Remark          string                       `fake:"{sentence}" json:"remark"`
}
