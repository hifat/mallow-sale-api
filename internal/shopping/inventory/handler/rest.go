package shoppingInventoryHandler

import (
	"github.com/gin-gonic/gin"
	shoppingInventoryModule "github.com/hifat/mallow-sale-api/internal/shopping/inventory"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	shoppingInventoryService shoppingInventoryModule.IService
}

func NewRest(shoppingInventoryService shoppingInventoryModule.IService) *Rest {
	return &Rest{shoppingInventoryService: shoppingInventoryService}
}

// @Summary 	Create ShoppingInventory
// @security 	BearerAuth
// @Tags 		Shopping Inventory
// @Accept 		json
// @Produce 	json
// @Param 		inventory body shoppingInventoryModule.Request true "Created inventory data"
// @Success 	201 {object} handling.ResponseItem[shoppingInventoryModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-inventories [post]
func (r *Rest) Create(c *gin.Context) {
	req := new(shoppingInventoryModule.Request)
	if err := c.ShouldBindJSON(req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.shoppingInventoryService.Create(c.Request.Context(), req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

// @Summary 	Find ShoppingInventory
// @security 	BearerAuth
// @Tags 		Shopping Inventory
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} handling.ResponseItems[shoppingInventoryModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-inventories [get]
func (r *Rest) Find(c *gin.Context) {
	res, err := r.shoppingInventoryService.Find(c.Request.Context())
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Shopping Inventory by ID
// @security 	BearerAuth
// @Tags 		Shopping Inventory
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Shopping Inventory ID"
// @Success 	200 {object} handling.Response
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-inventories/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.shoppingInventoryService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
