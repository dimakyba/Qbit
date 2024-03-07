package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)

	arr := extractDigits(k)

	PrintArray(arr)
}

func extractDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int

	for n != 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n /= 10
	}

	return digits
}

func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
