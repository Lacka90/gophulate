package computations

import "strconv"

// Fibonacci - fibonacci finder
func Fibonacci(n int) string {
	return strconv.Itoa(fibonacciInner(n))
}

func fibonacciInner(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciInner(n-1) + fibonacciInner(n-2)
}
