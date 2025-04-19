package usageUnit

import "github.com/hifat/cost-calculator-api/internal/entity"

type (
	UsageUnit struct {
		entity.Base `bson:"inline"`
		Code        string `bson:"code"`
		Name        string `bson:"name"`
	}

	UsageUnitEmbed struct {
		Code string `bson:"code"`
		Name string `bson:"name"`
	}
)
