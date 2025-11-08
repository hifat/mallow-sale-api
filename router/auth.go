package router

import (
	"github.com/gin-gonic/gin"
	authDi "github.com/hifat/mallow-sale-api/internal/auth/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database) {
	handler := authDi.Init(cfg, db)

	r.Group("/auth").
		POST("/signin", handler.Rest.Signin)

}
