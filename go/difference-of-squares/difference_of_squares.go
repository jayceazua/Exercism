package diffsquares

// SquareOfSum calculates square of sum of the first n numbers
func SquareOfSum(n int) int {
	sum := (n + 1) * n / 2
	return sum * sum
}

// SumOfSquares calculates sum of squares of the first n numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2 * n + 1) / 6
}

// Difference calculates the difference
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}