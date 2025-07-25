package settingsHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	settingModule "github.com/hifat/mallow-sale-api/internal/settings"
	settingService "github.com/hifat/mallow-sale-api/internal/settings/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	service settingService.Service
}

func NewRest(service settingService.Service) *Rest {
	return &Rest{service: service}
}

// @Summary      Update Settings
// @Tags         settings
// @Accept       json
// @Produce      json
// @Param        settings body settingModule.Request true "Settings data"
// @Success      200 {object} handling.SuccessResponse
// @Failure      400 {object} handling.ErrorResponse
// @Failure      500 {object} handling.ErrorResponse
// @Router       /settings [put]
func (h *Rest) Update(c *gin.Context) {
	var req settingModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	if err := h.service.Update(req.CostPercentage); err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, nil)
}

// @Summary      Get Settings
// @Tags         settings
// @Accept       json
// @Produce      json
// @Success      200 {object} settingModule.Entity
// @Failure      500 {object} handling.ErrorResponse
// @Router       /settings [get]
func (h *Rest) Get(c *gin.Context) {
	settings, err := h.service.Get()
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	c.JSON(http.StatusOK, settings)
}
