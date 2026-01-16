package shoppingHandler

import (
	"github.com/gin-gonic/gin"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type InventoryRest struct {
	service shoppingModule.IInventoryService
}

func NewInventoryRest(service shoppingModule.IInventoryService) *InventoryRest {
	return &InventoryRest{service: service}
}

// @Summary 	Create Shopping Inventory
// @security 	BearerAuth
// @Tags 		Shopping Inventory
// @Accept 		json
// @Produce 	json
// @Param 		inventory body shoppingModule.RequestShoppingInventory true "Created inventory data"
// @Success 	201 {object} handling.ResponseItem[shoppingModule.RequestShoppingInventory]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-inventories [post]
func (r *InventoryRest) Create(c *gin.Context) {
	req := new(shoppingModule.RequestShoppingInventory)
	if err := c.ShouldBindJSON(req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.service.Create(c.Request.Context(), req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

// @Summary 	Find Shopping Inventory
// @security 	BearerAuth
// @Tags 		Shopping Inventory
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} handling.ResponseItems[shoppingModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-inventories [get]
func (r *InventoryRest) Find(c *gin.Context) {
	res, err := r.service.Find(c.Request.Context())
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
func (r *InventoryRest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.service.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
