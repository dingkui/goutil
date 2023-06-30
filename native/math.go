package native

import (
	"math"
)

type m byte

const MathUtil m = iota

func (x m) Round(a float64, digits int) float64 {
	q := math.Pow(10, float64(digits))
	return math.Round(a*q) / q
}
