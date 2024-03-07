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

	var sum int
	for _, v := range input {
		sum += v
	}
	fmt.Println(sum)
	max := GetMaxIndex(input)
	min := GetMinIndex(input)

	fmt.Printf("%d %d\n", input[max], max+1)
	fmt.Printf("%d %d\n", input[min], min+1)
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

func GetMaxIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	maxIndex := 0
	maxValue := arr[0]

	for i, v := range arr {
		if v >= maxValue {
			maxIndex = i
			maxValue = v
		}
	}

	return maxIndex
}

func GetMinIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	minIndex := 0
	minValue := arr[0]

	for i, v := range arr {
		if v < minValue {
			minIndex = i
			minValue = v
		}
	}

	return minIndex
}
