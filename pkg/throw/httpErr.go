package throw

import (
	"errors"
	"net/http"

	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/pkg/utils/response"
)

var ErrRecordNotFound = errors.New("record not found")
var ErrInternalServer = errors.New(http.StatusText(http.StatusInternalServerError))

var CodeBadRequest = "BAD_REQUEST"
var CodeInternalServer = "INTERNAL_SERVER_ERROR"
var CodeRecordNotFound = "RECORD_NOT_FOUND"
var CodeInvalidForm = "INVALID_FROM"

var MsgInvalidForm = "invalid from"

func ValidateErr(err error) error {
	obj := response.ResponseErr{
		Response: response.Response{
			Status:  http.StatusBadRequest,
			Code:    CodeInvalidForm,
			Message: err.Error(),
		},
	}

	if attr, ok := err.(rules.ValidateErrs); ok {
		obj.Message = MsgInvalidForm
		obj.Attribute = attr
	}

	return obj
}

func BadRequestErr(err error) error {
	return response.ResponseErr{
		Response: response.Response{
			Status:  http.StatusBadRequest,
			Code:    CodeBadRequest,
			Message: err.Error(),
		},
	}
}

// If the error is not record not found, it will return 500.
func WhenRecordNotFoundErr(err error) error {
	if errors.Is(err, ErrRecordNotFound) {
		return response.ResponseErr{
			Response: response.Response{
				Status:  http.StatusNotFound,
				Code:    CodeRecordNotFound,
				Message: ErrRecordNotFound.Error(),
			},
		}
	}

	return InternalServerErr(err)
}

func InternalServerErr(err error) error {
	return response.ResponseErr{
		Response: response.Response{
			Status:  http.StatusInternalServerError,
			Code:    CodeInternalServer,
			Message: ErrInternalServer.Error(),
		},
	}
}

func Err(status int, code string, message string) error {
	if message == "" {
		message = http.StatusText(status)
	}

	return response.ResponseErr{
		Response: response.Response{
			Status:  status,
			Code:    code,
			Message: message,
		},
	}
}
