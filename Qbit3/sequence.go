package main

import (
	"fmt"
	"math"
)

func main() {
	var m, n uint
	fmt.Scan(&m, &n)

	if m <= n {
		for i := m; i <= n; i++ {
			a := math.Pow(-1.0, float64(i-1)) * (3*float64(i) - 1)
			fmt.Printf("%d ", int(a))
		}
	} else {
		for i := m; i >= n; i-- {
			a := math.Pow(-1.0, float64(i-1)) * (3*float64(i) - 1)
			fmt.Printf("%d ", int(a))
		}
	}

	fmt.Println()
}
