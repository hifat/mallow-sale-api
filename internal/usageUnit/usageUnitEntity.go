package usageUnit

import "github.com/hifat/mallow-sale-api/internal/entity"

type UsageUnit struct {
	entity.Base `bson:"inline"`
	Code        string `bson:"code"`
	Name        string `bson:"name"`
}

func (e *UsageUnit) DocName() string {
	return "usage_units"
}

type UsageUnitEmbed struct {
	Code string `bson:"code"`
	Name string `bson:"name"`
}

func (e *UsageUnitEmbed) SetAttr(code, name string) {
	e.Code = code
	e.Name = name
}
