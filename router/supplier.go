package router

import (
	"github.com/gin-gonic/gin"
	supplierDi "github.com/hifat/mallow-sale-api/internal/supplier/di"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func SupplierRouter(r *gin.RouterGroup, db *mongo.Database, logger logger.Logger) {
	handler := supplierDi.Init(db, logger)

	r.Group("/suppliers").
		GET("", handler.Rest.Find).
		GET(":id", handler.Rest.FindByID).
		POST("", handler.Rest.Create).
		PUT(":id", handler.Rest.UpdateByID).
		DELETE(":id", handler.Rest.DeleteByID)
}
