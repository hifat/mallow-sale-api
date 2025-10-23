package inventoryHandler

import (
	"github.com/gin-gonic/gin"
	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryService "github.com/hifat/mallow-sale-api/internal/inventory/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	inventoryService inventoryService.IService
}

func NewRest(inventoryService inventoryService.IService) *Rest {
	return &Rest{inventoryService: inventoryService}
}

// @Summary 	Create Inventory
// @Tags 		inventory
// @Accept 		json
// @Produce 	json
// @Param 		inventory body inventoryModule.Request true "Created inventory data"
// @Success 	201 {object} handling.ResponseItem[inventoryModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/inventories [post]
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

// @Summary 	Find Inventory by ID
// @Tags 		inventory
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "inventoryID"
// @Success 	200 {object} handling.ResponseItem[inventoryModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/inventories/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.inventoryService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Inventories
// @Tags 		inventory
// @Accept 		json
// @Produce 	json
// @Param 		query query utilsModule.QueryReq false "Query parameters"
// @Success 	200 {object} handling.ResponseItems[inventoryModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/inventories [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.inventoryService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Inventory by ID
// @Tags 		inventory
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "inventory ID"
// @Param 		inventory body inventoryModule.Request true "Updated inventory data"
// @Success 	200 {object} handling.ResponseItem[inventoryModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/inventories/{id} [put]
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

// @Summary 	Delete Inventory by ID
// @Tags 		inventory
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Inventory ID"
// @Success 	200 {object} handling.Response
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/inventories/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.inventoryService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
