package middlewareHandler

import (
	"github.com/gin-gonic/gin"
	middlewareService "github.com/hifat/mallow-sale-api/internal/middleware/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	middlewareService middlewareService.IService
}

func NewRest(middlewareService middlewareService.IService) *Rest {
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
