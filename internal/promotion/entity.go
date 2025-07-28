package promotionModule

import (
	recipeModule "github.com/hifat/mallow-sale-api/internal/recipe"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PromotionType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Code string             `bson:"code" json:"code"` // DISCOUNT, PAIR, FORCE_PRICE, OTHER
	Name string             `bson:"name" json:"name"`
}

type Entity struct {
	utilsModule.Base `bson:"inline"`

	Type     PromotionType         `bson:"type" json:"type"`
	Name     string                `bson:"name" json:"name"`
	Detail   string                `bson:"detail" json:"detail"`
	Discount float32               `bson:"discount" json:"discount"`
	Price    float32               `bson:"price" json:"price"`
	Products []recipeModule.Entity `bson:"products" json:"products"`
}
