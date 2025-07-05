package handling

type MetaResponse struct {
	Total int64 `json:"total"`
}

type ResponseItems[T any] struct {
	Items []T          `json:"items"`
	Meta  MetaResponse `json:"meta"`
}

type ResponseItem[T comparable] struct {
	Item T `json:"item"`
}
