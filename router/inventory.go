package router

import (
	"github.com/gin-gonic/gin"
	inventoryDi "github.com/hifat/mallow-sale-api/internal/inventory/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func InventoryRouter(r *gin.Engine, cfg *config.Config, db *mongo.Database) {
	handler := inventoryDi.Init(cfg, db)

	r.Group("/inventories").
		POST("/", handler.Rest.Create).
		GET("/:id", handler.Rest.FindByID).
		GET("/", handler.Rest.Find).
		PUT("/:id", handler.Rest.UpdateByID).
		DELETE("/:id", handler.Rest.DeleteByID)
}
