package promotionModule

import recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"

type ProtoType struct {
	ID        string                  `json:"id"`
	Type      PromotionTypeResponse   `json:"type"`
	Name      string                  `json:"name"`
	Detail    string                  `json:"detail"`
	Discount  float32                 `json:"discount"`
	Price     float32                 `json:"price"`
	Products  []recipeModule.Response `json:"products"`
	CreatedAt string                  `json:"createdAt"`
	UpdatedAt string                  `json:"updatedAt"`
}

func (r *ProtoType) GetProductIDs() []string {
	productIDs := make([]string, 0, len(r.Products))
	for _, product := range r.Products {
		productIDs = append(productIDs, product.ID)
	}

	return productIDs
}
