package usageUnitModule

type UsageUnitReq struct {
	Code string `validate:"required" json:"code"`

	Name string `json:"-"`
}

type Prototype struct {
	Code string `fake:"{lettern:5}" json:"code"`
	Name string `fake:"{name}" json:"name"`
}
