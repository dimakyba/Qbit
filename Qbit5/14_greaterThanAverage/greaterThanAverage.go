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

	// var sum int
	// for
	average := float32(sum(input)) / float32(len(input))

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

func sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}
