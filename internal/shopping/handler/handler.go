package shoppingHandler

type Handler struct {
	Rest          *Rest
	InventoryRest *InventoryRest
	ReceiptRest   *ReceiptRest
	UsageUnitRest *UsageUnitRest
}

func New(Rest *Rest, InventoryRest *InventoryRest, ReceiptRest *ReceiptRest, UsageUnitRest *UsageUnitRest) *Handler {
	return &Handler{
		Rest,
		InventoryRest,
		ReceiptRest,
		UsageUnitRest,
	}
}
