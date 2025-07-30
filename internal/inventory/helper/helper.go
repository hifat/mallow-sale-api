package helper

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
)

type Helper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error)
	IncressStock(ctx context.Context, id string, purchaseQuantity float32, purchasePrice float32) error
	DecressStock(ctx context.Context, id string, purchaseQuantity float32, purchasePrice float32) error
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

func (h *helper) currentPurchasePrice(inventory inventoryModule.Response, purchasePrice float32) float32 {
	remainingPricePerUnit := inventory.PurchasePrice / inventory.PurchaseQuantity
	remainingPrice := remainingPricePerUnit * inventory.PurchaseQuantity
	currentPrice := purchasePrice + remainingPrice

	return currentPrice
}

func (h *helper) IncressStock(ctx context.Context, id string, purchaseQuantity float32, purchasePrice float32) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, purchasePrice)
	currentQuantity := inventory.PurchaseQuantity + purchaseQuantity

	return h.inventoryRepository.IncressStock(ctx, id, currentQuantity, currentPrice)
}

func (h *helper) DecressStock(ctx context.Context, id string, purchaseQuantity float32, purchasePrice float32) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, purchasePrice)
	currentQuantity := inventory.PurchaseQuantity - purchaseQuantity

	return h.inventoryRepository.DecressStock(ctx, id, currentQuantity, currentPrice)
}
