package authModule

import (
	"context"
)

type IService interface {
	Signin(ctx context.Context, req *SigninReq) (*Passport, error)
	SigninWithGoogle(ctx context.Context, req *SigninReq) (*Passport, error)
	Auth(ctx context.Context, req *SigninReq) (*Passport, error)
}
