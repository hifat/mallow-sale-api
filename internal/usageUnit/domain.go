package usageUnitModule

type UsageUnitReq struct {
	Code string `validate:"required" json:"code"`

	Name string `json:"-"`
}
