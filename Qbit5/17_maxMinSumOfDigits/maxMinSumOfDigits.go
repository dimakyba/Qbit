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
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	// PrintArray(input)
	sumsOfDigits := GetSumsOfDigits(inputCopy)
	// PrintArray(sumsOfDigits)
	// PrintArray(input)
	max, min := GetMaxAndMin(input, sumsOfDigits)
	fmt.Printf("%d\n%d\n", input[min], input[max])
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

func GetSumsOfDigits(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		sum, n := 0, arr[i]
		for n > 0 {
			digit := n % 10
			sum += digit
			n /= 10
		}
		arr[i] = sum
	}
	return arr
}

func GetMaxAndMin(input []int, arr []int) (maxIndex, minIndex int) {
	if len(arr) == 0 {
		return -1, -1
	}

	maxIndex, minIndex = 0, 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] || (arr[i] == arr[maxIndex] && input[i] > input[maxIndex]) {
			maxIndex = i
		}
		if arr[i] < arr[minIndex] || (arr[i] == arr[minIndex] && input[i] < input[minIndex]) {
			minIndex = i
		}
	}

	return maxIndex, minIndex
}

// func PrintArray(arr []int) {
// 	for _, v := range arr {
// 		fmt.Printf("%d ", v)
// 	}
// 	fmt.Println()
// }
