package inventoryHandler

type Handler struct {
	InventoryRest *inventoryRest
}

func New(InventoryRest *inventoryRest) Handler {
	return Handler{
		InventoryRest,
	}
}
