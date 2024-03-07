package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 1.0

	for i := 1; i <= n; i++ {
		result *= 1 + float64(i)/10.0
	}

	fmt.Println(result)
}
