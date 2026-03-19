package settingModule

type Entity struct {
	CostPercentage float32 `bson:"cost_percentage" json:"costPercentage"`
	LinemanGP      float64 `bson:"lineman_gp" json:"linemanGP"`
	GrabGP         float64 `bson:"grab_gp" json:"grabGP"`
}
