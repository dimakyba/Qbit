package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var count, counter int
	var sum float64

	fmt.Scanln(&count)

	arr := make([]int, count)

	for i := 0; i < count; i++ {
		fmt.Scan(&arr[i])
		if isPrime(int(math.Abs(float64(arr[i])))) {
			counter++
			sum += float64(arr[i])
		}
	}

	if counter != 0 {
		fmt.Printf("%.3f\n", sum/float64(counter))
	} else {
		fmt.Println("0")
	}
}
