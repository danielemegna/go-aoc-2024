package day17

import "math"

func AdvOperation(numerator int, operator int) int {
	return int(float64(numerator) / math.Pow(float64(2), float64(operator)))
}
