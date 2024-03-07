package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var k int
	fmt.Scan(&k)
	input := getInputSlice()
	// input := []int{10, 2, 4, 5, 0, 2, 6, 10, -1, 0, -2}
	FindNextAfterMin(input)

}

func FindNextAfterMin(array []int) {
	min := math.MaxInt32

	for i, num := range array {
		if num < min {
			min = num
			var minIndex int = i
		}
	}

	nextAfterMin := math.MaxInt32
	nextAfterMinIndex := -1

	for i, num := range array {
		if num > min && num < nextAfterMin {
			nextAfterMin = num
			nextAfterMinIndex = i
		}
	}

	fmt.Printf("%d %d\n", nextAfterMin, nextAfterMinIndex+1)
}

func parseNumbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func getInputSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return parseNumbers(scanner.Text())
}
