package shoppingHandler

import (
	"github.com/gin-gonic/gin"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type UsageUnitRest struct {
	service shoppingModule.IUsageUnitService
}

func NewUsageUnitRest(service shoppingModule.IUsageUnitService) *UsageUnitRest {
	return &UsageUnitRest{service: service}
}

// @Summary 	Create Usage Unit
// @security 	BearerAuth
// @Tags 		Usage Unit
// @Accept 		json
// @Produce 	json
// @Param 		usageUnit body shoppingModule.RequestUsageUnit true "Created usage unit data"
// @Success 	201 {object} handling.ResponseItem[shoppingModule.RequestUsageUnit]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shopping-usage-units [post]
func (r *UsageUnitRest) Create(c *gin.Context) {
	req := new(shoppingModule.RequestUsageUnit)
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

// @Summary		Find Usage Unit
// @security	BearerAuth
// @Tags		Usage Unit
// @Accept		json
// @Produce		json
// @Success		200 {object} handling.ResponseItems[shoppingModule.ResUsageUnit]
// @Failure		400 {object} handling.ErrorResponse
// @Failure		404 {object} handling.ErrorResponse
// @Failure		500 {object} handling.ErrorResponse
// @Router		/shopping-usage-units [get]
func (r *UsageUnitRest) Find(c *gin.Context) {
	res, err := r.service.Find(c.Request.Context())
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary		Find Usage Unit by ID
// @security	BearerAuth
// @Tags		Usage Unit
// @Accept		json
// @Produce		json
// @Param		id path string true "Usage Unit ID"
// @Success		200 {object} handling.ResponseItem[shoppingModule.ResUsageUnit]
// @Failure		400 {object} handling.ErrorResponse
// @Failure		404 {object} handling.ErrorResponse
// @Failure		500 {object} handling.ErrorResponse
// @Router		/shopping-usage-units/{id} [get]
func (r *UsageUnitRest) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, err := r.service.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary		Update Usage Unit by ID
// @security	BearerAuth
// @Tags		Usage Unit
// @Accept		json
// @Produce		json
// @Param		id path string true "Usage Unit ID"
// @Param		usageUnit body shoppingModule.RequestUsageUnit true "Updated usage unit data"
// @Success		200 {object} handling.ResponseItem[shoppingModule.RequestUsageUnit]
// @Failure		400 {object} handling.ErrorResponse
// @Failure		404 {object} handling.ErrorResponse
// @Failure		500 {object} handling.ErrorResponse
// @Router		/shopping-usage-units/{id} [put]
func (r *UsageUnitRest) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	req := new(shoppingModule.RequestUsageUnit)
	if err := c.ShouldBindJSON(req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.service.UpdateByID(c.Request.Context(), id, req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary		Delete Usage Unit by ID
// @security		BearerAuth
// @Tags		Usage Unit
// @Accept		json
// @Produce		json
// @Param		id path string true "Usage Unit ID"
// @Success		200 {object} handling.ResponseItem[shoppingModule.RequestUsageUnit]
// @Failure		400 {object} handling.ErrorResponse
// @Failure		404 {object} handling.ErrorResponse
// @Failure		500 {object} handling.ErrorResponse
// @Router		/shopping-usage-units/{id} [delete]
func (r *UsageUnitRest) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := r.service.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, nil)
}
