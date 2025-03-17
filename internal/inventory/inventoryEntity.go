package inventory

import "github.com/hifat/cost-calculator-api/internal/entity"

type (
	InventoryEntity struct {
		entity.Base      `bson:"inline"`
		Name             string    `bson:"name"`              // วัตถุดิบ
		PurchasePrice    float64   `bson:"purchase_price"`    // ราคาซื้อ
		PurchaseQuantity float64   `bson:"purchase_quantity"` // ปริมาณซื้อ
		PurchaseUnit     UsageUnit `bson:"purchase_unit"`     // หน่วยซื้อ
		YieldPercentage  float64   `bson:"yield_percentage"`  // Yield %
		ActualPrice      float64   `bson:"actual_price"`      // ราคาจริง
		CostPerUnit      float64   `bson:"cost_per_unit"`     // ต้นทุน/หน่วย
		UsageQuantity    float64   `bson:"usage_quantity"`    // ปริมาณที่ใช้
		UsageUnit        UsageUnit `bson:"usage_unit"`        // หน่วยใช้
		UsageCost        float64   `bson:"usage_cost"`        // ต้นทุนที่ใช้
		Remark           string    `bson:"remark"`            // หมายเหตุ
	}

	UsageUnit struct {
		ID   string `bson:"id,omitempty"`
		Code string `bson:"code,omitempty"`
		Name string `bson:"name,omitempty"`
	}
)

func (m *InventoryEntity) DocName() string {
	return "inventories"
}
