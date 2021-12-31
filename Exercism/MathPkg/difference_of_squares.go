package diffsquares

import "math"

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = (i * (i + 1) / 2)
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	squares := 0
	for i := 1; i <= n; i++ {
		squares = (n * (n + 1) * ((2 * n) + 1)) / 6
	}
	return squares
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
