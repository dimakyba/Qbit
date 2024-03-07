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
	var average = (float32(input[max]) + float32(input[min])) / 2.0

	var counter int
	for _, v := range input {
		if float32(v) > average {
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
