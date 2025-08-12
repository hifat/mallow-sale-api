package utils

import "math"

// It returns an empty slice if the input is nil.
func MustToSlice[T any](v []T) []T {
	if v == nil {
		return []T{}
	}

	return v
}

func RoundToDecimals(value float64, decimals int) float64 {
	return math.Round(value*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
