package main

import (
	"fmt"
	"math"
)

func main() {
	var max, second_max, k int

	max = math.MinInt
	second_max = math.MinInt
	fmt.Scan(&k)
	for k != 0 {
		if k == max {
			second_max = k
		}
		if k > max {
			second_max = max
			max = k
		}
		if k > second_max && k < max {
			second_max = k
		}
		fmt.Scan(&k)
	}
	if second_max != math.MinInt {
		fmt.Println(second_max)
	}
}
