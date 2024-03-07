package main

import (
	"bufio"
	"fmt"
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

	counter := GetCounterOfNegativeBetweenTwo(input, l, r)

	fmt.Println(counter)
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

func GetCounterOfNegativeBetweenTwo(arr []int, l int, r int) int {
	if len(arr) == 0 {
		return -1
	}

	var counter int
	// minIndex := -1

	for i := l; i <= r; i++ {
		if arr[i] < 0 {
			// minIndex = i
			counter++
		}
	}

	return counter
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
