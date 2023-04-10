package fileutil

import (
	"math"
)

func Round1(a float64, digits int) float64 {
	q := math.Pow(10, float64(digits))
	return math.Round(a*q) / q
}
