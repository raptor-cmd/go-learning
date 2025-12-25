package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(7))

	var fibonacci func(numeroBase int) int

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	// 0, 1, 1, 2, 3, 5, 8, 13
	fmt.Println(fibonacci(7))
}
