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
	l := GetFirstMinIndex(input)
	r := GetLastMaxIndex(input)
	var counter int
	if l > r {
		l, r = r, l
	}
	for i := l; i <= r; i++ {
		if input[i] < 0 {
			counter++
		}
	}

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

func GetFirstMinIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	min := math.MaxInt64
	minIndex := -1

	for i, v := range arr {
		if v < min {
			min = v
			minIndex = i
		}
	}

	return minIndex
}

func GetLastMaxIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := math.MinInt64
	maxIndex := -1

	for i, v := range arr {
		if v >= max {
			max = v
			maxIndex = i
		}
	}

	return maxIndex
}
