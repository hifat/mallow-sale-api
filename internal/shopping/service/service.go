package shoppingService

import (
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	shoppingRepository "github.com/hifat/mallow-sale-api/internal/shopping/repository"
	"github.com/hifat/mallow-sale-api/pkg/logger"
)

type IService interface {
	Create(req *shoppingModule.Request) error
	UpdateIsComplete(req *shoppingModule.UpdateIsComplete) error
	Delete(id string) error
}

type service struct {
	logger        logger.ILogger
	shoppingRepo  shoppingRepository.IRepository
	inventoryRepo inventoryRepository.IRepository
}

func New(logger logger.ILogger, shoppingRepo shoppingRepository.IRepository, inventoryRepo inventoryRepository.IRepository) IService {
	return &service{
		logger,
		shoppingRepo,
		inventoryRepo,
	}
}

func (s *service) Create(req *shoppingModule.Request) error {
	return nil
}

func (s *service) UpdateIsComplete(req *shoppingModule.UpdateIsComplete) error {
	return nil
}

func (s *service) Delete(id string) error {
	return nil
}
