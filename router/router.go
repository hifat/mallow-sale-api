package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middlewareDi "github.com/hifat/mallow-sale-api/internal/middleware/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterAll(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) {
	AuthRouter(r, cfg, db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	m := middlewareDi.Init(cfg, db)
	r.Use(m.Rest.AuthGuard)

	InventoryRouter(r, cfg, db)
	RecipeRouter(r, cfg, db)
	SettingRouter(r, db)
	SupplierRouter(r, cfg, db)
	StockRouter(r, cfg, db)
	PricePreset(r, cfg, db)
	PromotionRouter(r, cfg, db)
	ShoppingRouter(r, cfg, db, grpcConn)
}
