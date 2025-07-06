package helper

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
)

type Helper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error)
}

type helper struct {
	inventoryRepository inventoryRepository.Repository
}

func New(inventoryRepository inventoryRepository.Repository) Helper {
	return &helper{
		inventoryRepository: inventoryRepository,
	}
}

func (h *helper) FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error) {
	inventories, err := h.inventoryRepository.FindInIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return func(id string) *inventoryModule.Response {
		for _, inventory := range inventories {
			if inventory.ID == id {
				return &inventory
			}
		}

		return nil
	}, nil
}
