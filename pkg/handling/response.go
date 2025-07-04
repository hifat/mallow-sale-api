package handling

type ResponseItems[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}

type ResponseItem[T comparable] struct {
	Item T `json:"item"`
}
