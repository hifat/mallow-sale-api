package settingModule

type Request struct {
	CostPercentage float32 `json:"costPercentage" validate:"required,gte=0,lte=100"`
}
