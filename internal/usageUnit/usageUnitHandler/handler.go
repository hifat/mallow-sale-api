package usageUnitHandler

type Handler struct {
	GRPC *usageUnitGRPC
}

func New(GRPC *usageUnitGRPC) Handler {
	return Handler{GRPC}
}
