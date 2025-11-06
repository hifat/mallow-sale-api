package authService

import (
	"context"

	authModule "github.com/hifat/mallow-sale-api/internal/auth"
	userRepository "github.com/hifat/mallow-sale-api/internal/user/repository"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Signin(ctx context.Context, req *authModule.SigninReq) (*authModule.AuthRes, error)
}

type service struct {
	logger   logger.ILogger
	userRepo userRepository.IRepository
}

func New(
	logger logger.ILogger,
	userRepo userRepository.IRepository,
) IService {
	return &service{
		logger,
		userRepo,
	}
}

func (s *service) Signin(ctx context.Context, req *authModule.SigninReq) (*authModule.AuthRes, error) {
	return nil, nil
}
