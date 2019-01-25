package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	sum := 0.0
	for i := range make([]int, n) {
		sum += float64(i + 1)
	}
	return int(math.Pow(sum, 2.0))
}

func SumOfSquares(n int) int {
	sum := 0
	for i := range make([]int, n) {
		sum += int(math.Pow(float64(i+1), 2.0))
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
