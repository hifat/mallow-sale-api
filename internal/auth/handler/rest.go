package authHandler

import (
	"github.com/gin-gonic/gin"
	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	authService "github.com/hifat/mallow-sale-api/internal/auth/service"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type Rest struct {
	authService authService.IService
}

func NewRest(authService authService.IService) *Rest {
	return &Rest{authService: authService}
}

// @Summary 	Signin
// @Tags 		auth
// @Accept 		json
// @Produce 	json
// @Param 		auth body authModule.SigninReq true "Sign"
// @Success 	200 {object} handling.ResponseItem[authModule.Passport]
// @Failure 	400 {object} handling.ErrorResponse
// @Failure 	404 {object} handling.ErrorResponse
// @Failure 	500 {object} handling.ErrorResponse
// @Router 		/auth/signin [post]
func (r *Rest) Signin(c *gin.Context) {
	var req authModule.SigninReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handling.ResponseFormErr(c, err)
		return
	}

	res, err := r.authService.Signin(c.Request.Context(), &req)
	if err != nil {
		handling.ResponseErr(c, err)
		return
	}

	handling.ResponseSuccess(c, *res)
}
