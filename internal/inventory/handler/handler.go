package inventoryHandler

type Handler struct {
	Rest *Rest
	GRPC *GRPC
}

func New(Rest *Rest, GRPC *GRPC) *Handler {
	return &Handler{
		Rest,
		GRPC,
	}
}
