package main

import "fmt"

func main() {
	var n int
	sum := 0
	fmt.Scan(&n)

	for n > 0 {
		sum += n & 1
		n >>= 1
	}
	fmt.Println(sum)
}
