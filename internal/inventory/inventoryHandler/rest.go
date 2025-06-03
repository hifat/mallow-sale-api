package inventoryHandler

import (
	core "github.com/hifat/goroger-core"
	"github.com/hifat/mallow-sale-api/internal/inventory"
	"github.com/hifat/mallow-sale-api/internal/inventory/inventoryService"
	"github.com/hifat/mallow-sale-api/pkg/utils/handlerUtils.go"
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
		handlerUtils.ResponseBadRequest(c, err)
	}

	if err := h.inventorySrv.Create(c.Context(), inventoryReq); err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseCreated(c)
}

func (h *inventoryRest) Find(c core.IHttpCtx) {
	res, err := h.inventorySrv.Find(c.Context())
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseItems(c, res)
}

func (h *inventoryRest) FindByID(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	res, err := h.inventorySrv.FindByID(c.Context(), inventoryID)
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseItem(c, res)
}

func (h *inventoryRest) Update(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	req := inventory.InventoryReq{}
	if err := c.BodyParser(&req); err != nil {
		handlerUtils.ResponseBadRequest(c, err)

		return
	}

	err := h.inventorySrv.Update(c.Context(), inventoryID, req)
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseCreated(c)
}

func (h *inventoryRest) Delete(c core.IHttpCtx) {
	inventoryID := c.Param("inventoryID")

	err := h.inventorySrv.Delete(c.Context(), inventoryID)
	if err != nil {
		handlerUtils.ResponseErr(c, err)

		return
	}

	handlerUtils.ResponseOK(c)
}
