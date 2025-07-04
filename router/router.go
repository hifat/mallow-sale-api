package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAll(r *gin.Engine, cfg *config.Config, db *mongo.Database) {
	InventoryRouter(r, cfg, db)
}
