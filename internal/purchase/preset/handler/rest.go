package purchasePresetHandler

import (
	"github.com/gin-gonic/gin"
	purchasePresetModule "github.com/hifat/mallow-sale-api/internal/purchase/preset"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	purchasePresetService purchasePresetModule.IService
}

func NewRest(purchasePresetService purchasePresetModule.IService) *Rest {
	return &Rest{purchasePresetService: purchasePresetService}
}

// @Summary 	Create Purchase Preset
// @security 	BearerAuth
// @Tags 		purchase-preset
// @Accept 		json
// @Produce 	json
// @Param 		purchase body purchasePresetModule.CreatePurchaseRequest true "Created purchase data"
// @Success 	201 {object} handling.ResponseItem[purchasePresetModule.CreatePurchaseRequest]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchase-presets [post]
func (r *Rest) Create(c *gin.Context) {
	var req purchasePresetModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchasePresetService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

// @Summary 	Find Purchases
// @security 	BearerAuth
// @Tags 		purchase-preset
// @Accept 		json
// @Produce 	json
// @Param 		query query utilsModule.QueryReq false "Query parameters"
// @Success 	200 {object} handling.ResponseItems[purchasePresetModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchase-presets [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchasePresetService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase-preset
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "purchase ID"
// @Success 	200 {object} handling.ResponseItem[purchasePresetModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchase-presets/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.purchasePresetService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase-preset
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "purchase ID"
// @Param 		purchase body purchasePresetModule.Request true "Updated purchase data"
// @Success 	200 {object} handling.ResponseItem[purchasePresetModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchase-presets/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req purchasePresetModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchasePresetService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase-preset
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Purchase ID"
// @Success 	200 {object} handling.ResponseItem[any]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchase-presets/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.purchasePresetService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, gin.H{"message": "Deleted successfully"})
}
