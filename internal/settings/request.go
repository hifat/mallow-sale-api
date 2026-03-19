package settingModule

type Request struct {
	CostPercentage float32 `fake:"{float32}" json:"costPercentage" validate:"required,gte=0,lte=100"`
	LinemanGP      float64 `fake:"{float64}" json:"linemanGP" validate:"required,gte=0,lte=100"`
	GrabGP         float64 `fake:"{float64}" json:"grabGP" validate:"required,gte=0,lte=100"`
}
