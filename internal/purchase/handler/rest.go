package purchaseHandler

import (
	"github.com/gin-gonic/gin"
	purchaseModule "github.com/hifat/mallow-sale-api/internal/purchase"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	purchaseService purchaseModule.IService
}

func NewRest(purchaseService purchaseModule.IService) *Rest {
	return &Rest{purchaseService: purchaseService}
}

// @Summary 	Create Purchase
// @security 	BearerAuth
// @Tags 		purchase
// @Accept 		json
// @Produce 	json
// @Param 		purchase body purchaseModule.CreatePurchaseRequest true "Created purchase data"
// @Success 	201 {object} handling.ResponseItem[purchaseModule.CreatePurchaseRequest]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchases [post]
func (r *Rest) Create(c *gin.Context) {
	var req purchaseModule.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchaseService.Create(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseCreated(c, *res)
}

// @Summary 	Find Purchases
// @security 	BearerAuth
// @Tags 		purchase
// @Accept 		json
// @Produce 	json
// @Param 		query query utilsModule.QueryReq false "Query parameters"
// @Success 	200 {object} handling.ResponseItems[purchaseModule.Response]
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchases [get]
func (r *Rest) Find(c *gin.Context) {
	var query utilsModule.QueryReq
	if err := c.ShouldBindQuery(&query); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchaseService.Find(c.Request.Context(), &query)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Find Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "purchase ID"
// @Success 	200 {object} handling.ResponseItem[purchaseModule.Response]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchases/{id} [get]
func (r *Rest) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := r.purchaseService.FindByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Update Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase
// @Accept 		json
// @Produce 	json
// @Param 		id path string true "purchase ID"
// @Param 		purchase body purchaseModule.CreatePurchaseRequest true "Updated purchase data"
// @Success 	200 {object} handling.ResponseItem[purchaseModule.CreatePurchaseRequest]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchases/{id} [put]
func (r *Rest) UpdateByID(c *gin.Context) {
	id := c.Param("id")

	var req purchaseModule.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.purchaseService.UpdateByID(c.Request.Context(), id, &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}

// @Summary 	Delete Purchase by ID
// @security 	BearerAuth
// @Tags 		purchase
// @Accept 		json
// @Produce 	json
// @Param 		id 	path string true "Purchase ID"
// @Success 	200 {object} handling.ResponseItem[any]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/purchases/{id} [delete]
func (r *Rest) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := r.purchaseService.DeleteByID(c.Request.Context(), id)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, gin.H{"message": "Deleted successfully"})
}
