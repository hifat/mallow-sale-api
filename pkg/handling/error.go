package handling

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/utils/token"
)

// TODO: Reflector, It duplicated prop
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	Status int `json:"-"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`

	Status int `json:"-"`
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
		define.CodeUnauthorized: {
			Code:    define.CodeUnauthorized,
			Message: define.MsgUnauthorized,
			Status:  http.StatusUnauthorized,
		},
		define.CodeInvalidCredentials: {
			Code:    define.CodeInvalidCredentials,
			Message: define.MsgInvalidCredentials,
			Status:  http.StatusUnauthorized,
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
		define.CodeDuplicatedInventoryName: {
			Code:    define.CodeDuplicatedInventoryName,
			Message: define.MsgDuplicatedInventoryName,
			Status:  http.StatusBadRequest,
		},
		define.CodeInvalidPurchaseUnit: {
			Code:    define.CodeInvalidPurchaseUnit,
			Message: define.MsgInvalidPurchaseUnit,
			Status:  http.StatusBadRequest,
		},
		define.CodeInvalidShoppingStatus: {
			Code:    define.CodeInvalidShoppingStatus,
			Message: define.MsgInvalidShoppingStatus,
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
	switch {
	case errors.Is(err, define.ErrRecordNotFound):
		return ErrorResponse{
			Code:    define.CodeRecordNotFound,
			Message: define.MsgRecordNotFound,
			Status:  http.StatusNotFound,
		}
	case errors.Is(err, token.ErrInvalidToken):
		return ErrorResponse{
			Code:    define.CodeInvalidToken,
			Message: define.MsgInvalidToken,
			Status:  http.StatusUnauthorized,
		}
	case errors.Is(err, token.ErrTokenExpired):
		return ErrorResponse{
			Code:    define.CodeTokenExpired,
			Message: define.MsgTokenExpired,
			Status:  http.StatusUnauthorized,
		}
	}

	return ErrorResponse{
		Code:    define.CodeInternalServerError,
		Message: define.MsgInternalServerError,
		Status:  http.StatusInternalServerError,
	}
}
