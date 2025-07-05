package handling

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hifat/mallow-sale-api/pkg/define"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	Status int `json:"-"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("code: %s, message: %s, status: %d", e.Code, e.Message, e.Status)
}

func ThrowErr(err error) ErrorResponse {
	if errors.Is(err, define.ErrRecordNotFound) {
		return ErrorResponse{
			Code:    define.CodeRecordNotFound,
			Message: define.MsgRecordNotFound,
			Status:  http.StatusNotFound,
		}
	}

	return ErrorResponse{
		Code:    define.CodeInternalServerError,
		Message: define.MsgInternalServerError,
		Status:  http.StatusInternalServerError,
	}
}
