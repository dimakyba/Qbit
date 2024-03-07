package main

import (
	"fmt"
)

func main() {
	var current_max, prev_max int
	counter := 0
	maxCounter := 0

	prev_max = -1

	for {
		fmt.Scan(&current_max)

		if current_max == 0 {
			break
		}

		if current_max == prev_max {
			counter++
		} else {
			if counter > maxCounter {
				maxCounter = counter
			}
			counter = 1
		}

		prev_max = current_max
	}

	if counter > maxCounter {
		maxCounter = counter
	}

	fmt.Println(maxCounter)
}
