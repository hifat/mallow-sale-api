package recipeModule_test

import (
	"reflect"
	"testing"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	usageUnitModule "github.com/hifat/mallow-sale-api/internal/usageUnit"
)

func Test_Request_GetUsageUnitCodes(t *testing.T) {
	tests := []struct {
		name string
		req  recipeModule.Request
		want []string
	}{
		{
			name: "should return usage unit codes from ingredients",
			req: recipeModule.Request{
				Ingredients: []recipeModule.IngredientRequest{
					{Unit: usageUnitModule.UsageUnitReq{Code: "KG"}},
					{Unit: usageUnitModule.UsageUnitReq{Code: "L"}},
					{Unit: usageUnitModule.UsageUnitReq{Code: "PCS"}},
				},
			},
			want: []string{"KG", "L", "PCS"},
		},
		{
			name: "should return empty slice when no ingredients",
			req: recipeModule.Request{
				Ingredients: []recipeModule.IngredientRequest{},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.req.GetUsageUnitCodes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsageUnitCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Request_GetInventoryIDs(t *testing.T) {
	tests := []struct {
		name string
		req  recipeModule.Request
		want []string
	}{
		{
			name: "should return inventory IDs from ingredients",
			req: recipeModule.Request{
				Ingredients: []recipeModule.IngredientRequest{
					{InventoryID: "1"},
					{InventoryID: "2"},
					{InventoryID: "3"},
				},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "should return empty slice when no ingredients",
			req: recipeModule.Request{
				Ingredients: []recipeModule.IngredientRequest{},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.req.GetInventoryIDs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventoryIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Prototype_GetInventoryIDs(t *testing.T) {
	tests := []struct {
		name string
		req  recipeModule.Prototype
		want []string
	}{
		{
			name: "should return inventoryIDs from ingredients",
			req: recipeModule.Prototype{
				Ingredients: []recipeModule.IngredientPrototype{
					{InventoryID: "1"},
					{InventoryID: "2"},
					{InventoryID: "3"},
				},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "should return empty slice when no ingredients",
			req: recipeModule.Prototype{
				Ingredients: []recipeModule.IngredientPrototype{},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.req.GetInventoryIDs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventoryIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Prototype_GetInventoryIDFromIngredients(t *testing.T) {
	tests := []struct {
		name string
		req  recipeModule.Prototype
		want []string
	}{
		{
			name: "should return inventory IDs from ingredients",
			req: recipeModule.Prototype{
				Ingredients: []recipeModule.IngredientPrototype{
					{Inventory: &inventoryModule.Prototype{ID: "1"}},
					{Inventory: &inventoryModule.Prototype{ID: "2"}},
					{Inventory: &inventoryModule.Prototype{ID: "3"}},
				},
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "should return empty slice when no ingredients",
			req: recipeModule.Prototype{
				Ingredients: []recipeModule.IngredientPrototype{},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.req.GetInventoryIDFromIngredients()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventoryIDFromIngredients() = %v, want %v", got, tt.want)
			}
		})
	}
}
