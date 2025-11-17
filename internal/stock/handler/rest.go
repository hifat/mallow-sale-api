package stockHandler

import (
	"github.com/gin-gonic/gin"
	stockModule "github.com/hifat/mallow-sale-api/internal/stock"
	stockService "github.com/hifat/mallow-sale-api/internal/stock/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	stockService stockService.IService
}

func NewRest(stockService stockService.IService) *Rest {
	return &Rest{stockService: stockService}
}

// @Summary     Create Stock
// @security 	BearerAuth
// @Tags        stock
// @Accept      json
// @Produce     json
// @Param       stock body stockModule.Request true "Created stock data"
// @Success     201 {object} handling.ResponseItem[stockModule.Request]
// @Failure     400 {object} handling.ErrorResponse
// @Failure     500 {object} handling.ErrorResponse
// @Router      /stocks [post]
func (r *Rest) Create(c *gin.Context) {
	var req stockModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	res, err := r.stockService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseCreated(c, *res)
}

// @Summary     Find Stocks
// @security 	BearerAuth
// @Tags        stock
// @Accept      json
// @Produce     json
// @Param       query query utilsModule.QueryReq false "Query parameters"
// @Success     200 {object} handling.ResponseItems[stockModule.Response]
// @Failure     500 {object} handling.ErrorResponse
// @Router      /stocks [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	res, err := r.stockService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, *res)
}

// @Summary     Find Stock by ID
// @security 	BearerAuth
// @Tags        stock
// @Accept      json
// @Produce     json
// @Param       id path string true "stockID"
// @Success     200 {object} handling.ResponseItem[stockModule.Response]
// @Failure     400 {object} handling.ErrorResponse
// @Failure     404 {object} handling.ErrorResponse
// @Failure     500 {object} handling.ErrorResponse
// @Router      /stocks/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")
	res, err := r.stockService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, *res)
}

// @Summary     Update Stock by ID
// @security 	BearerAuth
// @Tags        stock
// @Accept      json
// @Produce     json
// @Param       id path string true "stockID"
// @Param       stock body stockModule.Request true "Updated stock data"
// @Success     200 {object} handling.ResponseItem[stockModule.Request]
// @Failure     400 {object} handling.ErrorResponse
// @Failure     404 {object} handling.ErrorResponse
// @Failure     500 {object} handling.ErrorResponse
// @Router      /stocks/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var req stockModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}
	res, err := r.stockService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, *res)
}

// @Summary     Delete Stock by ID
// @security 	BearerAuth
// @Tags        stock
// @Accept      json
// @Produce     json
// @Param       id path string true "stockID"
// @Success     200 {object} handling.Response
// @Failure     404 {object} handling.ErrorResponse
// @Failure     500 {object} handling.ErrorResponse
// @Router      /stocks/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	err := r.stockService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}
	handling.ResponseSuccess(c, handling.Response{Message: "Deleted successfully"})
}
