package promotionModule

import (
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
)

type PromotionTypeRequest struct {
	ID   string `json:"id"`
	Code string `validate:"required" json:"code"` // DISCOUNT, PAIR, FORCE_PRICE, OTHER
	Name string `validate:"required" json:"name"`
}

type Request struct {
	Type     PromotionTypeRequest   `validate:"required" json:"type"`
	Name     string                 `validate:"required" json:"name"`
	Detail   string                 `json:"detail"`
	Discount float32                `json:"discount"`
	Price    float32                `json:"price"`
	Products []recipeModule.Request `json:"products"`
}

func (r *Request) Validate() error {
	switch r.Type.Code {
	case "DISCOUNT":
		if r.Discount <= 0 {
			return &ValidationError{Field: "discount", Message: "Discount is required for DISCOUNT type"}
		}
	case "PAIR":
		if len(r.Products) == 0 {
			return &ValidationError{Field: "product", Message: "At least one product is required for PAIR type"}
		}
	case "FORCE_PRICE":
		if r.Price <= 0 {
			return &ValidationError{Field: "price", Message: "Price is required for FORCE_PRICE type"}
		}
	default:
		return &ValidationError{Field: "type.code", Message: "Invalid promotion type"}
	}
	return nil
}

type PromotionTypeResponse struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

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

type Response struct {
	ProtoType
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
