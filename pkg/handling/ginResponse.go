package handling

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/mallow-sale-api/pkg/define"
)

func ResponseFormErr(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
		Code:    define.CodeInvalidForm,
		Message: err.Error(),
		Status:  http.StatusBadRequest,
	})
}

func ResponseErr(c *gin.Context, err error) {
	resErr, ok := err.(ErrorResponse)
	if !ok {
		resErr = ErrorResponse{
			Code:    define.CodeInternalServerError,
			Message: define.MsgInternalServerError,
			Status:  http.StatusInternalServerError,
		}
		return
	}

	c.AbortWithStatusJSON(resErr.Status, ErrorResponse{
		Code:    resErr.Code,
		Message: resErr.Message,
		Status:  resErr.Status,
	})
}

func ResponseCreated[T comparable](c *gin.Context, res ResponseItem[T]) {
	c.JSON(http.StatusCreated, res)
}

func ResponseCreatedBatch[T comparable](c *gin.Context, res ResponseItems[T]) {
	c.JSON(http.StatusCreated, res)
}

func ResponseSuccess(c *gin.Context, res any) {
	c.JSON(http.StatusOK, res)
}
