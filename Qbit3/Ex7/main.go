package main

import (
	"fmt"
	"math"
)

func main() {
	var current_max, prev_max int
	prev_max = math.MaxInt
	counter := 0
	current_max = 1

	for {
		fmt.Scan(&current_max)

		if current_max == 0 {
			break
		}
		if current_max > prev_max {
			counter++
		}
		prev_max = current_max
	}

	fmt.Println(counter)
}
