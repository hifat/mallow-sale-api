package promotionModule_test

import (
	"testing"

	promotionModule "github.com/hifat/mallow-sale-api/internal/promotion"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
)

func Test_GetProductIDs(t *testing.T) {
	protoType := &promotionModule.ProtoType{
		Products: []recipeModule.Response{
			{Prototype: recipeModule.Prototype{ID: "1"}},
			{Prototype: recipeModule.Prototype{ID: "2"}},
			{Prototype: recipeModule.Prototype{ID: "3"}},
		},
	}

	productIDs := protoType.GetProductIDs()
	if len(productIDs) != 3 {
		t.Errorf("expected 3 product IDs, got %d", len(productIDs))
	}
}
