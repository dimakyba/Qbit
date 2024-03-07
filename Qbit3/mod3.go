package main

import "fmt"

func main() {
	var n int
	k := 1
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if k%3 == 1 {
			fmt.Printf("%d ", k)
			k += 3
		}
	}
	fmt.Println()
}
