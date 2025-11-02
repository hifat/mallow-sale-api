package shoppingHandler

import (
	"github.com/gin-gonic/gin"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	shoppingService "github.com/hifat/mallow-sale-api/internal/shopping/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type ReceiptRest struct {
	receiptService shoppingService.IReceiptService
}

func NewReceiptRest(receiptService shoppingService.IReceiptService) *ReceiptRest {
	return &ReceiptRest{receiptService}
}

// @Summary 	Receipt Reader
// @Tags 		shopping
// @Accept 		json
// @Produce 	json
// @Param 		shopping body shoppingModule.ReqReceiptReader true "Receipt reader data"
// @Success 	201 {object} handling.ResponseItem[shoppingModule.ReqReceiptReader]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/shoppings/read-receipt [post]
func (r *ReceiptRest) Reader(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	req := shoppingModule.ReqReceiptReader{
		Image: image,
	}

	res, err := r.receiptService.Reader(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}
