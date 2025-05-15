package handlerUtils

import (
	"net/http"

	"github.com/hifat/cost-calculator-api/pkg/throw"
	"github.com/hifat/cost-calculator-api/pkg/utils/response"
	core "github.com/hifat/goroger-core"
)

func ResponseErr(c core.IHttpCtx, err error) {
	v, ok := err.(response.ResponseErr)
	if !ok {
		c.AbortWithJSON(http.StatusInternalServerError, throw.InternalServerErr(err))
		return
	}

	c.AbortWithJSON(v.Status, v)
}

func ResponseBadRequest(c core.IHttpCtx, err error) {
	c.AbortWithJSON(http.StatusBadRequest, throw.InternalServerErr(err))
}

func ResponseCreated(c core.IHttpCtx) {
	c.JSON(http.StatusCreated, response.Response{
		Code:    response.CodeCreated,
		Message: http.StatusText(http.StatusCreated),
	})
}

func ResponseOK(c core.IHttpCtx) {
	c.JSON(http.StatusOK, response.Response{
		Code:    response.CodeOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func ResponseItem[T comparable](c core.IHttpCtx, item T) {
	c.JSON(http.StatusOK, response.ResponseSuccess{
		Item: item,
	})
}

func ResponseItems[T any](c core.IHttpCtx, items []T) {
	c.JSON(http.StatusOK, response.ResponseSuccess{
		Items: items,
	})
}
