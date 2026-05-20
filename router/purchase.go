package router

import (
	"github.com/gin-gonic/gin"
	purchaseDi "github.com/hifat/mallow-sale-api/internal/purchase/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func PurchaseRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := purchaseDi.Init(cfg, db)

	r.Group("/purchases").
		GET("", handler.Rest.Find).
		GET(":id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT(":id", handler.Rest.UpdateByID).
		DELETE(":id", handler.Rest.DeleteByID)
}
