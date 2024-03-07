package main

import (
	"fmt"
)

func main() {
	var k int
	counter := 0
	fmt.Scan(&k)
	for k != 0 {
		if k%2 != 0 {
			counter++
		}
		fmt.Scan(&k)
	}
	fmt.Println(counter)
}
