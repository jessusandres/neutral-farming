package pkg

import (
	"math"
)

type Number interface {
	float64 | int | uint | int64
}

func RoundToDecimals(n float64, decimals uint32) float64 {
	scale := math.Pow(10, float64(decimals))

	return math.Round(n*scale) / scale
}

func AveragePercentForTwoValues[T Number](current, previous T) float64 {
	value := ((current - previous) / previous) * 100

	return RoundToDecimals(float64(value), 2)
}
