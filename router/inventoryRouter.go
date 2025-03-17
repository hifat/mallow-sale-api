package router

import inventoryDI "github.com/hifat/cost-calculator-api/internal/inventory/inventoryDi"

func (r *router) InventoryRouter() {
	handler := inventoryDI.InitInventory(r.cfg, r.db, r.logger)

	inventory := r.route.Group("/inventories")
	inventory.Get("", handler.InventoryRest.Find)
	inventory.Post("", handler.InventoryRest.Create)
}
