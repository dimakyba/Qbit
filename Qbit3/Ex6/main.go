package main

import (
	"fmt"
)

func main() {
	var i, k int
	counter := 1
	fmt.Scan(&i)

	for {
		fmt.Scan(&k)
		if k >= i {
			counter++
		}
		if i > k {
			break
		}
		i = k
	}

	fmt.Println(counter)
}
