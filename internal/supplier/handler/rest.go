package supplierHandler

import (
	"github.com/gin-gonic/gin"
	supplierModule "github.com/hifat/mallow-sale-api/internal/supplier"
	supplierService "github.com/hifat/mallow-sale-api/internal/supplier/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	supplierService supplierService.IService
}

func NewRest(supplierService supplierService.IService) *Rest {
	return &Rest{supplierService: supplierService}
}

// @Summary      Create Supplier
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        supplier body supplierModule.Request true "Created supplier data"
// @Success      201 {object} handling.ResponseItem[supplierModule.Request]
// @Failure      400 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /suppliers [post]
func (r *Rest) Create(c *gin.Context) {
	var req supplierModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	res, err := r.supplierService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseCreated(c, *res)
}

// @Summary      Find Suppliers
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        query query utilsModule.QueryReq false "Query parameters"
// @Success      200 {object} handling.ResponseItems[supplierModule.Response]
// @Failure      500 {object} handling.ErrorResponse
// @Router       /suppliers [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.supplierService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary      Find Supplier by ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id path string true "supplierID"
// @Success      200 {object} handling.ResponseItem[supplierModule.Response]
// @Failure      400 {object} handling.ErrorResponse
// @Failure      404 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /suppliers/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, err := r.supplierService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary      Update Supplier by ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id path string true "supplierID"
// @Param        supplier body supplierModule.Request true "Updated supplier data"
// @Success      200 {object} handling.ResponseItem[supplierModule.Request]
// @Failure      400 {object} handling.ErrorResponse
// @Failure      404 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /suppliers/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var req supplierModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.supplierService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary      Delete Supplier by ID
// @Tags         supplier
// @Accept       json
// @Produce      json
// @Param        id path string true "supplierID"
// @Success      200 {object} handling.SuccessResponse
// @Failure      404 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /suppliers/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := r.supplierService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, handling.SuccessResponse{Message: "Deleted successfully"})
}
