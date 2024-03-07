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
	max, min := GetMaxAndMin(input)
	var sum int

	for _, v := range input {
		sum += v
	}
	sum -= input[min] + input[max]

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

func GetMaxAndMin(arr []int) (maxIndex, minIndex int) {
	if len(arr) == 0 {
		return -1, -1
	}

	maxIndex, minIndex = 0, 0

	for i, v := range arr {
		if v > arr[maxIndex] {
			maxIndex = i
		}
		if v < arr[minIndex] {
			minIndex = i
		}
	}

	return maxIndex, minIndex
}
