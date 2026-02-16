package middlewareService

import (
	"context"
	"strings"

	middlewareModule "github.com/hifat/mallow-sale-api/internal/middleware"
	userModule "github.com/hifat/mallow-sale-api/internal/user"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/handling"
	"github.com/hifat/mallow-sale-api/pkg/logger"
	"github.com/hifat/mallow-sale-api/pkg/utils/token"
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
) middlewareModule.IService {
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
