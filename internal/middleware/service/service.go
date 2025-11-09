package middlewareService

import (
	"context"
	"strings"

	userRepository "github.com/hifat/mallow-sale-api/internal/user/repository"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"github.com/hifat/mallow-sale-api/pkg/utils/token"
)

type IService interface {
	AuthGuard(ctx context.Context, t string) error
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

func (s *service) AuthGuard(ctx context.Context, t string) error {
	tk := strings.TrimPrefix(t, "Bearer ")
	_, err := token.Claims(s.cfg.Auth, token.ACCESS, tk)
	if err != nil {
		s.logger.Error(err)
		return handling.ThrowErr(err)
	}

	return nil
}
