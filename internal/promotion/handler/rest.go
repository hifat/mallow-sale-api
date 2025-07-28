package promotionHandler

import (
	"github.com/gin-gonic/gin"
	promotionModule "github.com/hifat/mallow-sale-api/internal/promotion"
	promotionService "github.com/hifat/mallow-sale-api/internal/promotion/service"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	promotionService promotionService.Service
}

func NewRest(promotionService promotionService.Service) *Rest {
	return &Rest{promotionService: promotionService}
}

// @Summary 	Create Promotion
// @Tags 		promotion
// @Accept 		json
// @Produce 	json
// @Param 		promotion body promotionModule.Request true "Created promotion data"
// @Success 	201 {object} handling.ResponseItem[promotionModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/promotions [post]
func (r *Rest) Create(c *gin.Context) {
	var req promotionModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.promotionService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Promotions
// @Tags 		promotion
// @Accept 		json
// @Produce 	json
// @Param 		query query utilsModule.QueryReq false "Query parameters"
// @Success 	200 {object} handling.ResponseItems[promotionModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/promotions [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.promotionService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Promotion by ID
// @Tags 		promotion
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "promotionID"
// @Success 	200 {object} handling.ResponseItem[promotionModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/promotions/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.promotionService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Promotion by ID
// @Tags 		promotion
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Promotion ID"
// @Param 		promotion body promotionModule.Request true "Updated promotion data"
// @Success 	200 {object} handling.ResponseItem[promotionModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/promotions/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req promotionModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.promotionService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Promotion by ID
// @Tags 		promotion
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Promotion ID"
// @Success 	200 {object} handling.SuccessResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/promotions/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	if err := r.promotionService.DeleteByID(c.Request.Context(), id); err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, handling.SuccessResponse{Message: "Promotion deleted successfully"})
}
