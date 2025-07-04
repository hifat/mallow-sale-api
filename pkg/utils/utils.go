package utils

// It returns an empty slice if the input is nil.
func MustToSlice[T any](v []T) []T {
	if v == nil {
		return []T{}
	}

	return v
}
