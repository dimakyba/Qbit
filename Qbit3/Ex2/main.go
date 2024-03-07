package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	counter := 1
	k := math.MinInt
	fmt.Scan(&n)

	for n != 0 {
		if n > k {
			k = n
			counter = 1
		} else if k == n {
			counter++
		}
		fmt.Scan(&n)
	}
	fmt.Println(counter)
}
