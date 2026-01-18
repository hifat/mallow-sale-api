package storageHandler

type Handler struct {
	Rest *Rest
}

func New(rest *Rest) *Handler {
	return &Handler{
		Rest: rest,
	}
}
