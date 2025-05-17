package throw

import (
	"errors"
	"net/http"

	"github.com/hifat/goroger-core/rules"
	"github.com/hifat/mallow-sale-api/pkg/utils/response"
)

var ErrRecordNotFound = errors.New("record not found")

var ErrBadRequestCode = "BAD_REQUEST"
var ErrInternalServerCode = "INTERNAL_SERVER_ERROR"
var ErrRecordNotFoundCode = "RECORD_NOT_FOUND"
var ErrInvalidFormCode = "INVALID_FROM"

func ValidateErr(err error) error {
	obj := response.ResponseErr{
		Response: response.Response{
			Status:  http.StatusBadRequest,
			Code:    ErrInvalidFormCode,
			Message: err.Error(),
		},
	}

	if attr, ok := err.(rules.ValidateErrs); ok {
		obj.Message = http.StatusText(http.StatusBadRequest)
		obj.Attribute = attr
	}

	return obj
}

func BadRequestErr(err error) error {
	return response.ResponseErr{
		Response: response.Response{
			Status:  http.StatusBadRequest,
			Code:    ErrBadRequestCode,
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
				Code:    ErrRecordNotFoundCode,
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
			Code:    ErrInternalServerCode,
			Message: http.StatusText(http.StatusInternalServerError),
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
