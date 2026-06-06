package router

import (
	"github.com/gin-gonic/gin"
	purchasePresetDi "github.com/hifat/mallow-sale-api/internal/purchase/preset/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func PurchasePresetRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := purchasePresetDi.Init(cfg, db)

	r.Group("/purchase-presets").
		GET("", handler.Rest.Find).
		GET(":id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT(":id", handler.Rest.UpdateByID).
		DELETE(":id", handler.Rest.DeleteByID)
}
