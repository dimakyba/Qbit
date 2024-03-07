package main

import (
	"fmt"
)

func main() {
	var n, sum, max int
	fmt.Scan(&n)

	max = n
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			sum += n
		} else {
			n = 3*n + 1
		}
		if n > max {
			max = n
		}
	}
	fmt.Printf("%d %d\n", sum, max)
}
