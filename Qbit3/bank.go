package main

import (
	"bufio"
	"fmt"
	"math"
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
	inputNumbers := GetInputSlice()
	var n, p, m, i float64
	n = float64(inputNumbers[0])
	p = float64(inputNumbers[1]) / 100.0
	m = float64(inputNumbers[2])

	for i = 1; n*math.Pow(1+p, float64(i)) <= m; i++ {
	}
	fmt.Println(i)
}
