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

	for _, v := range input {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
		}
	}

	for _, v := range input {
		if v%2 != 0 {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
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
