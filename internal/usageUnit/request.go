package usageUnitModule

type ReqUsageUnit struct {
	Code string `validate:"required" json:"code"`

	Name string `json:"-"`
}
