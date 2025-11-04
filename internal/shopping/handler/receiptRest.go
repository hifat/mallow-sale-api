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
// @Description  **Error Codes:**
// @Description
// @Description  **400 Bad Request:**
// @Description  - `MAX_FILE_SIZE`: max file size
// @Description  - `NOT_ALLOWED_MIME_TYPE`: not allowed mime type
// @Description
// @Description  **500 Internal Server Error:**
// @Description  - `INTERNAL_SERVER_ERROR`: internal server error
// @Param 		shopping body shoppingModule.ReqReceiptReader true "Receipt reader data"
// @Success 	200 {object} handling.ResponseItems[shoppingModule.ResReceiptReader]
// @Failure 	400 {object} handling.ErrorResponse
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
