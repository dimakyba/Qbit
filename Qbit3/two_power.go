package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k float64
	fmt.Scan(&n)
	k = 0
	for math.Pow(2, k) < n {
		k++
	}
	fmt.Println(k)
}
