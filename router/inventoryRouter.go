package router

import inventoryDI "github.com/hifat/cost-calculator-api/internal/inventory/inventoryDi"

func (r *router) InventoryRouter() {
	handler := inventoryDI.Init(r.cfg, r.db, r.logger)

	inventory := r.route.Group("/api/inventories")
	inventory.Get("", handler.InventoryRest.Find)
	inventory.Get("/:inventoryID", handler.InventoryRest.FindByID)
	inventory.Post("", handler.InventoryRest.Create)
	inventory.Put("/:inventoryID", handler.InventoryRest.Update)
	inventory.Delete("/:inventoryID", handler.InventoryRest.Delete)
}
