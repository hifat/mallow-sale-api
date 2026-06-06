package purchasePresetHandler

type Handler struct {
	Rest *Rest
}

func New(Rest *Rest) *Handler {
	return &Handler{
		Rest,
	}
}
