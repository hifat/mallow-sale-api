package usageUnitModule

type UsageUnitReq struct {
	Code string `validate:"required" json:"code"`

	Name string `json:"-"`
}

type Prototype struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
