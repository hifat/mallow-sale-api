package router

import (
	"github.com/gin-gonic/gin"
	stockDi "github.com/hifat/mallow-sale-api/internal/stock/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func StockRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := stockDi.Init(cfg, db)

	r.Group("/stocks").
		GET("", handler.Rest.Find).
		GET(":id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT(":id", handler.Rest.UpdateByID).
		DELETE(":id", handler.Rest.DeleteByID)
}
