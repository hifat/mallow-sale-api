package promotionModule

type PromotionTypeResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Response struct {
	ProtoType
}
