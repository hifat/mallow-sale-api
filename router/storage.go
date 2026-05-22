package router

import (
	"github.com/gin-gonic/gin"
	storageDi "github.com/hifat/mallow-sale-api/internal/storage/di"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func StorageRouter(r *gin.RouterGroup, cfg *config.Config, db *mongo.Database, grpcConn *grpc.ClientConn) {
	handler := storageDi.Init(cfg, db, grpcConn)

	r.Group("/uploads").
		POST("", handler.Rest.Upload)
}
