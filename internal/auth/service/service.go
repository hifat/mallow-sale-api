package authService

import (
	"context"
	"errors"

	"cloud.google.com/go/auth/credentials/idtoken"
	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"github.com/hifat/mallow-sale-api/pkg/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	logger   logger.ILogger
	cfg      *config.Config
	userRepo userModule.IRepository
}

func New(
	logger logger.ILogger,
	cfg *config.Config,
	userRepo userModule.IRepository,
) authModule.IService {
	return &service{
		logger,
		cfg,
		userRepo,
	}
}

func (s *service) Signin(ctx context.Context, req *authModule.SigninReq) (*authModule.Passport, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(define.ErrRecordNotFound, err) {
			return nil, handling.ThrowErrByCode(define.CodeInvalidCredentials)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidCredentials)
	}

	passport := &authModule.Passport{
		User: &authModule.AuthRes{
			Prototype: userModule.Prototype{
				Name:     user.Name,
				Username: user.Username,
			},
		},
	}

	t := token.New(s.cfg, *passport)

	_, accessToken, err := t.Signed(token.ACCESS)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	passport.SetAccessToken(accessToken)

	_, refreshToken, err := t.Signed(token.REFRESH)
	passport.SetRefreshToken(refreshToken)

	return passport, nil
}

func (s *service) SigninWithGoogle(ctx context.Context, req *authModule.SigninReq) (*authModule.Passport, error) {
	payload, err := idtoken.Validate(context.Background(), req.Token, s.cfg.Google.ClientID)
	if err != nil {
		return nil, handling.ThrowErrByCode(define.CodeInvalidCredentials)
	}

	email, _ := payload.Claims["email"].(string)

	user, err := s.userRepo.FindByUsername(ctx, email)
	if err != nil {
		if errors.Is(define.ErrRecordNotFound, err) {
			return nil, handling.ThrowErrByCode(define.CodeInvalidCredentials)
		}

		s.logger.Error(err)
		return nil, handling.ThrowErr(err)
	}

	passport := &authModule.Passport{
		User: &authModule.AuthRes{
			Prototype: userModule.Prototype{
				Name:     user.Name,
				Username: user.Username,
			},
		},
	}

	t := token.New(s.cfg, *passport)

	_, accessToken, err := t.Signed(token.ACCESS)
	if err != nil {
		return nil, handling.ThrowErr(err)
	}

	passport.SetAccessToken(accessToken)

	_, refreshToken, err := t.Signed(token.REFRESH)
	passport.SetRefreshToken(refreshToken)

	return passport, nil
}

func (s *service) Auth(ctx context.Context, req *authModule.SigninReq) (*authModule.Passport, error) {
	switch req.LoginType {
	case authModule.EnumInternal:
		return s.Signin(ctx, req)
	case authModule.EnumGoogle:
		return s.SigninWithGoogle(ctx, req)
	default:
		return nil, handling.ThrowErrByCode(define.CodeInvalidLoginType)
	}
}
