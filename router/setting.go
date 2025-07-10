package router

import (
	"github.com/gin-gonic/gin"
	settingDi "github.com/hifat/mallow-sale-api/internal/settings/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func SettingRouter(r *gin.RouterGroup, db *mongo.Database) {
	h := settingDi.InitializeSettingsRest(db)

	r.GET("/settings", h.Get)
	r.PUT("/settings", h.Update)
}
