package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n*2; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
