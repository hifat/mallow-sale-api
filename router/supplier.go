package router

import (
	"github.com/gin-gonic/gin"
	supplierDi "github.com/hifat/mallow-sale-api/internal/supplier/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func SupplierRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := supplierDi.Init(cfg, db)

	r.Group("/suppliers").
		GET("", handler.Rest.Find).
		GET(":id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT(":id", handler.Rest.UpdateByID).
		DELETE(":id", handler.Rest.DeleteByID)
}
