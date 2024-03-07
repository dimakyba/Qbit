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
	input := GetInputSlice()

	lens := len(input)

	var sum int
	for i, v := range input {
		sum += v * int(math.Pow(10, float64(lens-i-1)))

	}

	sum += 42
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
