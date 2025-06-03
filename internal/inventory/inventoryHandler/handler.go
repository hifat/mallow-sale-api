package inventoryHandler

type Handler struct {
	InventoryRest *inventoryRest
	GRPC          *inventoryGRPC
}

func New(InventoryRest *inventoryRest, GRPC *inventoryGRPC) Handler {
	return Handler{
		InventoryRest,
		GRPC,
	}
}
