package router

import (
	"github.com/gin-gonic/gin"
	shoppingDi "github.com/hifat/mallow-sale-api/internal/shopping/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func ShoppingRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) {
	h := shoppingDi.Init(cfg, db, grpcConn)

	r.Group("/shoppings").
		GET("", h.Rest.Find).
		GET("/:id", h.Rest.FindByID).
		POST("", h.Rest.Create).
		POST("/read-receipt", h.ReceiptRest.Reader).
		PUT("/:id", h.Rest.UpdateByID).
		PATCH("/:id/status", h.Rest.UpdateStatus).
		DELETE("/:id", h.Rest.DeleteByID)

	r.Group("/shopping-inventories").
		GET("", h.InventoryRest.Find).
		POST("", h.InventoryRest.Create).
		DELETE("/:id", h.InventoryRest.DeleteByID)

}
