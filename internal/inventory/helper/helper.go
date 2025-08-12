package helper

import (
	"context"
	"math"

	inventoryModule "github.com/hifat/mallow-sale-api/internal/inventory"
	inventoryRepository "github.com/hifat/mallow-sale-api/internal/inventory/repository"
	"github.com/hifat/mallow-sale-api/pkg/utils"
)

type Helper interface {
	FindAndGetByID(ctx context.Context, ids []string) (func(id string) *inventoryModule.Response, error)
	IncreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float64, purchasePrice float64) error
	DecreaseStock(ctx context.Context, inventoryID string, purchaseQuantity float64, purchasePrice float64) error
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

func (h *helper) currentPurchasePrice(inventory inventoryModule.Response, reqPurchasePrice float64, isIncrease bool) float64 {
	const epsilon = 1e-9 // ค่า tolerance ที่เล็กมาก สำหรับ float64

	// ตรวจสอบกรณีพิเศษก่อน
	if inventory.PurchaseQuantity == 0 {
		if isIncrease {
			return utils.RoundToDecimals(reqPurchasePrice, 3)
		}
		return 0
	}

	// คำนวณ unit price
	remainingPricePerUnit := inventory.PurchasePrice / inventory.PurchaseQuantity

	// คำนวณ total remaining price
	// Round ที่นี่เพื่อป้องกัน accumulating error
	remainingPrice := utils.RoundToDecimals(remainingPricePerUnit*inventory.PurchaseQuantity, 3)

	var currentPrice float64
	if isIncrease {
		currentPrice = reqPurchasePrice + remainingPrice
	} else {
		currentPrice = remainingPrice - reqPurchasePrice
	}

	// ถ้าผลลัพธ์ใกล้ 0 มาก (น้อยกว่า epsilon) ให้ return 0
	if math.Abs(currentPrice) < epsilon {
		return 0
	}

	return utils.RoundToDecimals(currentPrice, 3)
}

func (h *helper) IncreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float64, reqPurchasePrice float64) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, reqPurchasePrice, true)
	currentQuantity := inventory.PurchaseQuantity + reqPurchaseQuantity

	return h.inventoryRepository.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}

func (h *helper) DecreaseStock(ctx context.Context, inventoryID string, reqPurchaseQuantity float64, reqPurchasePrice float64) error {
	inventory, err := h.inventoryRepository.FindByID(ctx, inventoryID)
	if err != nil {
		return err
	}

	currentPrice := h.currentPurchasePrice(*inventory, reqPurchasePrice, false)
	currentQuantity := inventory.PurchaseQuantity - reqPurchaseQuantity

	return h.inventoryRepository.UpdateStock(ctx, inventoryID, currentQuantity, currentPrice)
}
