package helper

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
)

type Helper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error)
	IncreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float32, purchasePrice float32) error
	DecreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float32, purchasePrice float32) error
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

func (h *helper) currentPurchasePrice(inventory inventoryModule.Response, reqPurchasePrice float32) float32 {
	remainingPricePerUnit := float32(0.0)
	if inventory.PurchasePrice != 0 || inventory.PurchaseQuantity != 0 {
		remainingPricePerUnit = inventory.PurchasePrice / inventory.PurchaseQuantity
	}

	remainingPrice := remainingPricePerUnit * inventory.PurchaseQuantity
	currentPrice := reqPurchasePrice + remainingPrice

	return currentPrice
}

func (h *helper) IncreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float32, reqPurchasePrice float32) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, reqPurchasePrice)
	currentQuantity := inventory.PurchaseQuantity + reqPurchaseQuantity

	return h.inventoryRepository.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}

func (h *helper) DecreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float32, reqPurchasePrice float32) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, reqPurchasePrice)
	currentQuantity := inventory.PurchaseQuantity - reqPurchaseQuantity

	return h.inventoryRepository.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}
