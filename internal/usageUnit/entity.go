package usageUnitModule

type Entity struct {
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}
