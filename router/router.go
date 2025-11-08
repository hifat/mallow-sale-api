package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterAll(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) {
	AuthRouter(r, cfg, db)
	InventoryRouter(r, cfg, db)
	RecipeRouter(r, cfg, db)
	SettingRouter(r, db)
	SupplierRouter(r, cfg, db)
	StockRouter(r, cfg, db)
	PromotionRouter(r, cfg, db)
	ShoppingRouter(r, cfg, db, grpcConn)
}
