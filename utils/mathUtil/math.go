package mathUtil

import (
	"math"
)

func Round(a float64, digits int) float64 {
	q := math.Pow(10, float64(digits))
	return math.Round(a*q) / q
}
