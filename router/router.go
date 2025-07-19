package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAll(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	InventoryRouter(r, cfg, db)
	RecipeRouter(r, cfg, db)
	SettingRouter(r, db)
	SupplierRouter(r, db, logger.New())
}
