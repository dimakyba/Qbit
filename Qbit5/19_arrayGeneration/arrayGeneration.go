package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	var a []int = make([]int, 10)
	a[0] = x

	for i := 1; i < len(a); i++ {
		a[i] = a[i-1]*a[i-1]%100 - 5*a[i-1] + 6
	}

	PrintArray(a)
}

func PrintArray(a []int) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
