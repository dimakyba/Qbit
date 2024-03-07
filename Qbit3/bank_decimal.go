package main

import (
	"fmt"
)

func main() {
	var N, n float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	// First loop
	N = 0
	for i := 1.0; N < n; i++ {
		N += 1 / i
		if N >= n {
			fmt.Printf("%.0f ", i-1)
			break
		}
	}

	// Reset N
	N = 0

	// Second loop
	for i := 1.0; N <= n; i++ {
		N += 1 / i
		if N > n {
			fmt.Printf("%.0f\n", i)
			break
		}
	}
}
