package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	for i := 0; i < k; i++ {
		fmt.Println(n)
	}
}
