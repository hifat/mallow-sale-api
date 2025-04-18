package usageUnit

import "github.com/hifat/cost-calculator-api/internal/entity"

type (
	UsageUnit struct {
		entity.Base `bson:"inline"`
		Code        string `bson:"code,omitempty"`
		Name        string `bson:"name,omitempty"`
	}

	UsageUnitEmbed struct {
		Code string `bson:"code,omitempty"`
		Name string `bson:"name,omitempty"`
	}
)
