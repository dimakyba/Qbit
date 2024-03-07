package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0

	for i := 1; count < n; i++ {
		for j := 1; j <= i; j++ {
			if count < n {
				fmt.Printf("%d ", i)
				count++
			} else {
				break
			}
		}
	}
	fmt.Println()
}
