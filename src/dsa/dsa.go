package dsa

func FibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FactorialRecursive(n int) int {
	if n == 1 {
		return 1
	}

	return n * FactorialRecursive(n-1)
}
