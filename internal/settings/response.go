package settingModule

type Response struct {
	CostPercentage float32 `fake:"{float32}" json:"costPercentage"`
	LinemanGP      float64 `fake:"{float64}" json:"linemanGP"`
	GrabGP         float64 `fake:"{float64}" json:"grabGP"`
}
