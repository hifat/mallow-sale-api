package inventoryHelper

import (
	"context"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	"github.com/hifat/mallow-sale-api/pkg/utils"
)

//go:generate mockgen -source=./helper.go -destination=./mock/helper.go -package=mockInventoryHelper
type IHelper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error)
	IncreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float64, purchasePrice float64) error
	DecreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float64, purchasePrice float64) error
}

type helper struct {
	inventoryRepo inventoryRepository.IRepository
}

func New(inventoryRepo inventoryRepository.IRepository) IHelper {
	return &helper{
		inventoryRepo: inventoryRepo,
	}
}

func (h *helper) FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error) {
	inventories, err := h.inventoryRepo.FindInIDs(ctx, ids)
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

func (h *helper) calPurchasePrice(inventory inventoryModule.Response, reqPurchasePrice float64, isIncrease bool) float64 {
	// Check special case before
	if inventory.PurchaseQuantity == 0 {
		if isIncrease {
			return utils.RoundToDecimals(reqPurchasePrice, 3)
		}

		return 0
	}

	invPurchasePrice := utils.RoundToDecimals(inventory.PurchasePrice, 3)

	var currentPrice float64
	if isIncrease {
		currentPrice = invPurchasePrice + reqPurchasePrice
	} else {
		currentPrice = invPurchasePrice - reqPurchasePrice
	}

	return utils.RoundToDecimals(currentPrice, 3)
}

func (h *helper) IncreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float64, reqPurchasePrice float64) error {
	inventory, err := h.inventoryRepo.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.calPurchasePrice(*inventory, reqPurchasePrice, true)
	currentQuantity := inventory.PurchaseQuantity + reqPurchaseQuantity

	return h.inventoryRepo.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}

func (h *helper) DecreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float64, reqPurchasePrice float64) error {
	inventory, err := h.inventoryRepo.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.calPurchasePrice(*inventory, reqPurchasePrice, false)
	currentQuantity := inventory.PurchaseQuantity - reqPurchaseQuantity

	return h.inventoryRepo.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}
