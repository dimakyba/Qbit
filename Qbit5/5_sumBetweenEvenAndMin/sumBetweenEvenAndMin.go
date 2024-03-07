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
	l := GetFirstEvenIndex(input)
	r := GetFirstMinIndex(input)
	// fmt.Printf("%d %d", l, r)
	var sum int
	if l == -1 || r == -1 {
		for _, v := range input {
			sum += v
		}
	} else {
		if l > r {
			l, r = r, l
		}
		for i := l; i <= r; i++ {
			sum += input[i]
		}
	}

	fmt.Println(sum)
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

func GetFirstEvenIndex(arr []int) (firstIndex int) {
	firstIndex = -1

	for i, v := range arr {
		if v%2 == 0 {
			if firstIndex == -1 {
				firstIndex = i
				return firstIndex
			}
		}
	}

	return
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
