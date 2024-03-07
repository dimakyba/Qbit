package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	k := math.MinInt
	fmt.Scan(&n)

	for n != 0 {
		if n > k {
			k = n
		}
		fmt.Scan(&n)
	}
	fmt.Println(k)
}
