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
	l, r := GetFirstAndLastEvenIndexes(input)
	var sum int
	if l > r {
		l, r = r, l
	}
	for i := l; i <= r; i++ {
		if input[i] > 0 {
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

func GetFirstAndLastEvenIndexes(arr []int) (firstIndex, lastIndex int) {
	firstIndex = -1
	lastIndex = -1

	for i, v := range arr {
		if v%2 == 0 {
			if firstIndex == -1 {
				firstIndex = i
			}
			lastIndex = i
		}
	}

	return firstIndex, lastIndex
}
