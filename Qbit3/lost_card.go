package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	inputNumbers := GetInputSlice()
	n := inputNumbers[0]
	inputNumbers = inputNumbers[1:]

	for _, num := range inputNumbers {
		sum += num
	}

	expectedSum := n * (n + 1) / 2
	lostValue := expectedSum - sum
	fmt.Println(lostValue)
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
