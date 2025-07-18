package router

import (
	"github.com/gin-gonic/gin"
	recipeDi "github.com/hifat/mallow-sale-api/internal/recipe/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func RecipeRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := recipeDi.Init(cfg, db)

	r.Group("/recipes").
		GET("", handler.Rest.Find).
		GET("/:id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT("/:id", handler.Rest.UpdateByID).
		DELETE("/:id", handler.Rest.DeleteByID).
		PATCH("/order-no", handler.Rest.PatchNoBatch)
}
