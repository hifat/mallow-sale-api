package settingModule

type Request struct {
	CostPercentage float32 `fake:"{float32}" json:"costPercentage" validate:"required,gte=0,lte=100"`
}

type Response struct {
	CostPercentage float32 `fake:"{float32}" json:"costPercentage"`
}
