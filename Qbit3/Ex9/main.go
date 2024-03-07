package main

import (
	"fmt"
)

func main() {
	var k int
	sum := 0

	for k != 0 {
		fmt.Scan(&k)
		sum += k

	}

	fmt.Println(sum)
}
