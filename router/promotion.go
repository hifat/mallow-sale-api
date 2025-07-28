package router

import (
	"github.com/gin-gonic/gin"
	promotionDi "github.com/hifat/mallow-sale-api/internal/promotion/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func PromotionRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := promotionDi.Init(cfg, db)

	r.Group("/promotions").
		GET("", handler.Rest.Find).
		GET("/:id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT("/:id", handler.Rest.UpdateByID).
		DELETE("/:id", handler.Rest.DeleteByID)
}
