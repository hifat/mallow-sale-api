package usageUnit

import "github.com/hifat/mallow-sale-api/internal/entity"

type UsageUnit struct {
	entity.Base `bson:"inline"`
	Code        string `bson:"code"`
	Name        string `bson:"name"`
}

func (m *UsageUnit) Doc() string {
	return "usage_units"
}

type UsageUnitEmbed struct {
	Code string `fake:"{name}" bson:"code"`
	Name string `fake:"{name}" bson:"name"`
}

func (e *UsageUnitEmbed) SetAttr(code, name string) {
	e.Code = code
	e.Name = name
}
