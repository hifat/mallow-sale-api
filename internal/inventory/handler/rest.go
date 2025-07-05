package inventoryHandler

import (
	"fmt"

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

// Create Inventory
// @Summary Create a new inventory item
// @Description Create a new inventory item with the provided details
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body inventoryModule.Request true "Inventory item to create"
// @Success 201 {object} inventoryModule.Response
// @Failure 400 {object} handling.ErrorResponse
// @Failure 500 {object} handling.ErrorResponse
// @Router /inventories [post]
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

// Find Inventory By ID
// @Summary Get inventory item by ID
// @Description Retrieve a specific inventory item by its ID
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "inventoryID"
// @Success 200 {object} inventoryModule.Response
// @Failure 400 {object} handling.ErrorResponse
// @Failure 404 {object} handling.ErrorResponse
// @Failure 500 {object} handling.ErrorResponse
// @Router /inventories/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.inventoryService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary Ping endpoint
// @Description A simple ping endpoint to check if the server is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Server is running"
// @Router /inventories/ping [get]
func (r *Rest) Ping(c *gin.Context) {}

// Find Inventories
// @Summary Get inventories
// @Description Retrieve all inventory items
// @Tags inventory
// @Accept json
// @Produce json
// @Success 200 {object} handling.ResponseItems[inventoryModule.Response]
// @Failure 500 {object} handling.ErrorResponse
// @Router /inventories [get]
func (r *Rest) Find(c *gin.Context) {
	res, err := r.inventoryService.Find(c.Request.Context())
	if err != nil {
		fmt.Println("err", err)
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// Update Inventory By ID
// @Summary Update inventory item by ID
// @Description Update a specific inventory item by its ID
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "inventory ID"
// @Param inventory body inventoryModule.Request true "Updated inventory data"
// @Success 200 {object} inventoryModule.Response
// @Failure 400 {object} handling.ErrorResponse
// @Failure 404 {object} handling.ErrorResponse
// @Failure 500 {object} handling.ErrorResponse
// @Router /inventories/{id} [put]
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

// Delete Inventory By ID
// @Summary Delete inventory item by ID
// @Description Delete a specific inventory item by its ID
// @Tags inventory
// @Accept json
// @Produce json
// @Param id path string true "Inventory item ID"
// @Success 200 {object} handling.SuccessResponse
// @Failure 400 {object} handling.ErrorResponse
// @Failure 404 {object} handling.ErrorResponse
// @Failure 500 {object} handling.ErrorResponse
// @Router /inventories/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.inventoryService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
