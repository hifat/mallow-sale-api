package router

import (
	"github.com/gin-gonic/gin"
	inventoryDi "github.com/hifat/mallow-sale-api/internal/inventory/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func InventoryRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := inventoryDi.Init(cfg, db)

	r.Group("/inventories").
		GET("", handler.Rest.Find).
		GET("/:id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT("/:id", handler.Rest.UpdateByID).
		DELETE("/:id", handler.Rest.DeleteByID)
}
