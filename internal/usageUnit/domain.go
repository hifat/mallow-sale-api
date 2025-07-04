package usageUnitModule

type UsageUnitReq struct {
	Code string `validate:"required" json:"code"`

	Name string `json:"-"`
}

type Prototype struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Entity struct {
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}
