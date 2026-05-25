package supplierInventoryHandler

import (
	"github.com/gin-gonic/gin"
	supplierInventoryModule "github.com/hifat/mallow-sale-api/internal/supplier/inventory"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	supplierInventoryService supplierInventoryModule.IService
}

func NewRest(supplierInventoryService supplierInventoryModule.IService) *Rest {
	return &Rest{supplierInventoryService: supplierInventoryService}
}

// @Summary     Find Inventories Grouped by SupplierID
// @security 	BearerAuth
// @Tags        Supplier Inventory
// @Accept      json
// @Produce     json
// @Param       query query utilsModule.QueryReq false "Query parameters"
// @Success     200 {object} handling.ResponseItems[supplierInventoryModule.GroupBySupplierResponse]
// @Failure     500 {object} handling.ErrorResponse
// @Router      /supplier-inventories [get]
func (r *Rest) FindGroupBySupplier(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.supplierInventoryService.FindGroupBySupplier(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}
