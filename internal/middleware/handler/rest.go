package middlewareHandler

import (
	"github.com/gin-gonic/gin"
	middlewareModule "github.com/hifat/mallow-sale-api/internal/middleware"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	middlewareService middlewareModule.IService
}

func NewRest(middlewareService middlewareModule.IService) *Rest {
	return &Rest{middlewareService: middlewareService}
}

func (r *Rest) AuthGuard(c *gin.Context) {
	t := c.GetHeader("Authorization")

	err := r.middlewareService.AuthGuard(c.Request.Context(), t)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	c.Next()
}
