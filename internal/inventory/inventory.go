package inventory

import (
	"time"

	"github.com/hifat/mallow-sale-api/internal/usageUnit"
)

type UsageUnitRes struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (s *UsageUnitRes) SetAttr(code, name string) {
	s.Code = code
	s.Name = name
}

type InventoryPrototype struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	PurchasePrice    float32    `json:"purchasePrice"`
	PurchaseQuantity float32    `json:"purchaseQuantity"`
	YieldPercentage  float32    `json:"yieldPercentage"`
	Remark           string     `json:"remark"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

type InventoryReq struct {
	Name             string  `fake:"{firstname}" validate:"required" json:"name"`
	PurchasePrice    float32 `fake:"{number}" validate:"required" json:"purchasePrice"`
	YieldPercentage  float32 `fake:"{number}" validate:"required" json:"yieldPercentage"`
	Remark           string  `fake:"{sentence}" validate:"required,max=255" json:"remark"`
	PurchaseQuantity float32 `fake:"{number}" validate:"required" json:"purchaseQuantity"`

	PurchaseUnitCode string                   ` fake:"ML" validate:"required" json:"purchaseUnitCode"`
	PurchaseUnit     usageUnit.UsageUnitEmbed `fake:"-" json:"-"`
}

type InventoryRes struct {
	InventoryPrototype
	PurchaseUnit UsageUnitRes `json:"purchaseUnit"`
}

type FilterReq struct {
	IDs   []string `json:"ids"`
	Codes []string `json:"codes"`
}
