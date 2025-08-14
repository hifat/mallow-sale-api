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

func getErrObject(code string) ErrorResponse {
	mapErr := map[string]ErrorResponse{
		define.CodeRecordNotFound: {
			Code:    define.CodeRecordNotFound,
			Message: define.MsgRecordNotFound,
			Status:  http.StatusNotFound,
		},
		define.CodeInvalidUsageUnit: {
			Code:    define.CodeInvalidUsageUnit,
			Message: define.MsgInvalidUsageUnit,
			Status:  http.StatusBadRequest,
		},
		define.CodeInvalidInventoryID: {
			Code:    define.CodeInvalidInventoryID,
			Message: define.MsgInvalidInventoryID,
			Status:  http.StatusBadRequest,
		},
		define.CodeOrderNoMustBeUnique: {
			Code:    define.CodeOrderNoMustBeUnique,
			Message: define.MsgOrderNoMustBeUnique,
			Status:  http.StatusBadRequest,
		},
		define.CodeInvalidSupplierID: {
			Code:    define.CodeInvalidSupplierID,
			Message: define.MsgInvalidSupplierID,
			Status:  http.StatusBadRequest,
		},
		define.CodeInvalidRecipeType: {
			Code:    define.CodeInvalidRecipeType,
			Message: define.MsgInvalidRecipeType,
			Status:  http.StatusBadRequest,
		},
		define.CodeInventoryNameAlreadyExists: {
			Code:    define.CodeInventoryNameAlreadyExists,
			Message: define.MsgInventoryNameAlreadyExists,
			Status:  http.StatusBadRequest,
		},
	}

	errObj, ok := mapErr[code]
	if !ok {
		return ErrorResponse{
			Code:    define.CodeInternalServerError,
			Message: define.MsgInternalServerError,
			Status:  http.StatusInternalServerError,
		}
	}

	return errObj
}

func ThrowErrByCode(code string) ErrorResponse {
	return getErrObject(code)
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
