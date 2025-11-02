package shoppingHandler

type Handler struct {
	Rest        *Rest
	ReceiptRest *ReceiptRest
}

func New(Rest *Rest, ReceiptRest *ReceiptRest) *Handler {
	return &Handler{
		Rest,
		ReceiptRest,
	}
}
