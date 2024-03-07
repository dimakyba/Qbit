package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	var n int
	fmt.Scan(&n)
	counter := 0
	inputNumbers := GetInputSlice()
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", inputNumbers[counter])
			counter++
		}
		fmt.Println()
	}
}
