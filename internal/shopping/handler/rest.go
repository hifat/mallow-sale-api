package shoppingHandler

import (
	"github.com/gin-gonic/gin"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	shoppingService shoppingModule.IService
}

func NewRest(shoppingService shoppingModule.IService) *Rest {
	return &Rest{shoppingService: shoppingService}
}

// @Summary 	Create Shopping
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		shopping body shoppingModule.Request true "Created shopping data"
// @Success 	201 {object} handling.ResponseItem[shoppingModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings [post]
func (r *Rest) Create(c *gin.Context) {
	var req shoppingModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.shoppingService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

// @Summary 	Create Shoppings
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		shoppings body []shoppingModule.Request true "Created shoppings data"
// @Success 	201 {object} handling.ResponseItems[shoppingModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/batch [post]
func (r *Rest) CreateBatch(c *gin.Context) {
	var reqs []*shoppingModule.Request
	if err := c.ShouldBindJSON(&reqs); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.shoppingService.CreateBatch(c.Request.Context(), reqs)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreatedBatch(c, *res)
}

// @Summary 	Find Shoppings
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} handling.ResponseItems[shoppingModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings [get]
func (r *Rest) Find(c *gin.Context) {
	res, err := r.shoppingService.Find(c.Request.Context())
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Shopping by ID
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Shopping ID"
// @Success 	200 {object} handling.ResponseItem[shoppingModule.Response]
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.shoppingService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Shopping
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "Shopping ID"
// @Param 		shopping body shoppingModule.Request true "Updated shopping data"
// @Success 	200 {object} handling.ResponseItem[shoppingModule.Request]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req shoppingModule.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.shoppingService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Shopping Status
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "shopping ID"
// @Param 		shopping body shoppingModule.ReqUpdateStatus true "Updated shopping status"
// @Success 	200 {object} handling.Response
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/{id}/status [patch]
func (r *Rest) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var req shoppingModule.ReqUpdateStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.shoppingService.UpdateStatus(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Shopping by ID
// @security 	BearerAuth
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Shopping ID"
// @Success 	200 {object} handling.Response
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.shoppingService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, res)
}
