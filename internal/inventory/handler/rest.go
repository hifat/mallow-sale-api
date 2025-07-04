package inventoryHandler

import (
	"github.com/gin-gonic/gin"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryService "github.com/hifat/mallow-sale-api/internal/inventory/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	inventoryService inventoryService.Service
}

func NewRest(inventoryService inventoryService.Service) *Rest {
	return &Rest{inventoryService: inventoryService}
}

func (r *Rest) Create(c *gin.Context) {
	var req inventoryModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.inventoryService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.inventoryService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

func (r *Rest) Find(c *gin.Context) {
	res, err := r.inventoryService.Find(c.Request.Context())
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req inventoryModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.inventoryService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.inventoryService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
