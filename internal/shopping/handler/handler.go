package shoppingHandler

type Handler struct {
	Rest          *Rest
	InventoryRest *InventoryRest
	ReceiptRest   *ReceiptRest
}

func New(Rest *Rest, InventoryRest *InventoryRest, ReceiptRest *ReceiptRest) *Handler {
	return &Handler{
		Rest,
		InventoryRest,
		ReceiptRest,
	}
}
