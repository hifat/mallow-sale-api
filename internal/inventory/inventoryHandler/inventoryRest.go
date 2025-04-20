package inventoryHandler

import (
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

	if err := h.inventorySrv.Create(c.Context(), inventoryReq); err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"message": "ok",
	})
}

func (h *inventoryRest) Find(c core.IHttpCtx) {
	res, err := h.inventorySrv.Find(c.Context())
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

func (h *inventoryRest) FindByID(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	res, err := h.inventorySrv.FindByID(c.Context(), inventoryID)
	if err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"item": res,
	})
}

func (h *inventoryRest) Update(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	req := inventory.InventoryReq{}
	if err := c.BodyParser(&req); err != nil {
		c.AbortWithJSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

		return
	}

	err := h.inventorySrv.Update(c.Context(), inventoryID, req)
	if err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "ok",
	})
}

func (h *inventoryRest) Delete(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	err := h.inventorySrv.Delete(c.Context(), inventoryID)
	if err != nil {
		c.AbortWithJSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "ok",
	})
}
