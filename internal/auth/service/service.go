package authService

import (
	"context"
	"errors"

	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
	userRepository "github.com/hifat/mallow-sale-api/internal/user/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"github.com/hifat/mallow-sale-api/pkg/utils/token"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Signin(ctx context.Context, req *authModule.SigninReq) (*authModule.Passport, error)
}

type service struct {
	logger   logger.ILogger
	cfg      *config.Config
	userRepo userRepository.IRepository
}

func New(
	logger logger.ILogger,
	cfg *config.Config,
	userRepo userRepository.IRepository,
) IService {
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
