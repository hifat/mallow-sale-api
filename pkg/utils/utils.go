package utils

import (
	"math"
)

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

type IID interface {
	GetID() string
}

func GetIDs[T IID](items []T) []string {
	ids := make([]string, 0, len(items))
	for _, v := range items {
		ids = append(ids, v.GetID())
	}

	return ids
}

type ICode interface {
	GetCode() string
}

func GetCodes[T ICode](items []T) []string {
	codes := make([]string, 0, len(items))
	for _, v := range items {
		codes = append(codes, v.GetCode())
	}

	return codes
}
