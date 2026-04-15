package router

import (
	"github.com/gin-gonic/gin"
	pricePresetDi "github.com/hifat/mallow-sale-api/internal/pricePreset/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func PricePreset(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := pricePresetDi.Init(cfg, db)

	r.Group("/price-presets").
		GET("", handler.Rest.Find).
		GET("/:id", handler.Rest.FindByID)

}
