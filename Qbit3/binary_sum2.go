package main

import "fmt"

func main() {
	var n int
	sum := 0
	fmt.Scan(&n)

	for n > 0 {
		if n%2 != 0 {
			sum++
		}
		n /= 2
	}
	fmt.Println(sum)
}
