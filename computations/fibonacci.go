package computations

// Fibonacci - fibonacci finder
func Fibonacci(n int) interface{} {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1).(int) + Fibonacci(n-2).(int)
}
