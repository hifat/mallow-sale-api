package router

import (
	"github.com/gin-gonic/gin"
	supplierInventoryDi "github.com/hifat/mallow-sale-api/internal/supplier/inventory/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func SupplierInventoryRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := supplierInventoryDi.Init(cfg, db)

	r.Group("/supplier-inventories").
		GET("", handler.Rest.FindGroupBySupplier)
}
