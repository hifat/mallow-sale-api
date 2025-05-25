package usageUnit

type UsageUnitProtoType struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type FilterReq struct {
	Codes []string `json:"codes"`
}
