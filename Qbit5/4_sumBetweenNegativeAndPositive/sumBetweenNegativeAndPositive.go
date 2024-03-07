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
	l, r := GetFirstNegativeAndPositive(input)
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

func GetFirstNegativeAndPositive(arr []int) (negativeIndex, positiveIndex int) {
	negativeIndex = -1
	positiveIndex = -1

	for i, v := range arr {
		if v < 0 && negativeIndex == -1 {
			negativeIndex = i
		} else if v > 0 && positiveIndex == -1 {
			positiveIndex = i
		}

		if negativeIndex != -1 && positiveIndex != -1 {
			break
		}
	}

	return negativeIndex, positiveIndex
}
