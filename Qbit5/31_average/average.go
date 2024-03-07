package main

import (
	"fmt"
)

func main() {

	var k int
	fmt.Scan(&k)
	var input []int = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&input[i])
	}
	result := CalculateAverageOfOddElements(input)
	fmt.Printf("%.3f\n", result)
}

func CalculateAverageOfOddElements(arr []int) float64 {
	var sum float64
	var counter float64

	for i := 1; i < len(arr); i += 2 {
		if arr[i]%2 != 0 {
			sum += float64(arr[i])
			fmt.Println(sum)
			counter++
		}
		if counter == 0 {
			return 0
		}

	}

	return sum / float64(counter)
}
