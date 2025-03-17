package inventoryHandler

import (
	"context"
	"net/http"

	"github.com/hifat/cost-calculator-api/internal/inventory"
	"github.com/hifat/cost-calculator-api/internal/inventory/inventoryService"
	core "github.com/hifat/goroger-core"
)

type inventoryRest struct {
	inventorySrv inventoryService.IInventoryService
}

func NewRest(inventorySrv inventoryService.IInventoryService) *inventoryRest {
	return &inventoryRest{inventorySrv}
}

func (h *inventoryRest) Create(c core.IHttpCtx) {
	inventoryReq := inventory.InventoryReq{}
	if err := c.BodyParser(&inventoryReq); err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if err := h.inventorySrv.Create(context.Background(), inventoryReq); err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "ok",
	})
}

func (h *inventoryRest) Find(c core.IHttpCtx) {
	res, err := h.inventorySrv.Find(context.Background())
	if err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"items": res,
	})
}
