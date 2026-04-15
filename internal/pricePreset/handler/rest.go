package pricePresetHandler

import (
	"github.com/gin-gonic/gin"
	pricePresetModule "github.com/hifat/mallow-sale-api/internal/pricePreset"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	pricePresetService pricePresetModule.IService
}

func NewRest(pricePresetService pricePresetModule.IService) *Rest {
	return &Rest{pricePresetService: pricePresetService}
}

// @Summary     Find Price Presets
// @security 	BearerAuth
// @Tags        price_preset
// @Accept      json
// @Produce     json
// @Param       query query utilsModule.QueryReq false "Query parameters"
// @Success     200 {object} handling.ResponseItems[pricePresetModule.Response]
// @Failure     500 {object} handling.ErrorResponse
// @Router      /price-presets [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	res, err := r.pricePresetService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, *res)
}

// @Summary     Find Price Preset by ID
// @security 	BearerAuth
// @Tags        price_preset
// @Accept      json
// @Produce     json
// @Param       id path string true "pricePresetID"
// @Success     200 {object} handling.ResponseItem[pricePresetModule.Response]
// @Failure     400 {object} handling.ErrorResponse
// @Failure     404 {object} handling.ErrorResponse
// @Failure     500 {object} handling.ErrorResponse
// @Router      /price-presets/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, err := r.pricePresetService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, *res)
}
