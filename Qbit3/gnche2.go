package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var number int
	fmt.Scan(&number)

	counter := 0
	a := 1
	var output []string

	for counter < number {
		repeatCount := int(math.Min(float64(number-counter), float64(a)))
		for i := 0; i < repeatCount; i++ {
			output = append(output, strconv.Itoa(a))
		}
		counter += a
		a++
	}

	fmt.Println(strings.Join(output, " "))
}
