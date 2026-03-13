package authModule

import (
	"context"
)

type IService interface {
	Signin(ctx context.Context, req *SigninReq) (*Passport, error)
}
