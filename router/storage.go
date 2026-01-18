package router

import (
	"github.com/gin-gonic/gin"
	storageDi "github.com/hifat/mallow-sale-api/internal/storage/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
)

func StorageRouter(r *gin.RouterGroup, cfg *config.Config) {
	handler, err := storageDi.Init(cfg)
	if err != nil {
		panic(err)
	}

	r.Group("/storage").
		POST("/upload", handler.Rest.Upload)
}
