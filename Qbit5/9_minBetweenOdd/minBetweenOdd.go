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
	input := GetInputSlice()
	l, r := GetFirstAndSecondOddIndexes(input)
	if l > r {
		l, r = r, l
	}

	min := GetMinBetweenTwo(input, l, r)

	fmt.Println(min)
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func GetInputSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return numbers(scanner.Text())
}

func GetMinBetweenTwo(arr []int, l int, r int) int {
	if len(arr) == 0 {
		return -1
	}

	min := math.MaxInt64
	// minIndex := -1

	for i := l; i <= r; i++ {
		if arr[i] < min {
			min = arr[i]
			// minIndex = i
		}
	}

	return min
}

func GetFirstAndSecondOddIndexes(arr []int) (firstIndex, secondIndex int) {
	firstIndex = -1
	secondIndex = -1

	for i, v := range arr {
		if v%2 != 0 {
			if firstIndex == -1 {
				firstIndex = i
			} else {
				secondIndex = i
				break
			}
		}
	}

	return firstIndex, secondIndex
}
