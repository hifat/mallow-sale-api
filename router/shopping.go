package router

import (
	"github.com/gin-gonic/gin"
	shoppingDi "github.com/hifat/mallow-sale-api/internal/shopping/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func ShoppingRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	h := shoppingDi.Init(cfg, db)

	r.Group("/shoppings").
		GET("", h.Rest.Find).
		POST("", h.Rest.Create).
		PATCH("/:id/is-complete", h.Rest.UpdateIsComplete).
		DELETE("/:id", h.Rest.DeleteByID)
}
