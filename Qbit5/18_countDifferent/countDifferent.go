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

	uniqueNumbers := countUniqueNumbers(input)
	fmt.Println(uniqueNumbers)
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
func countUniqueNumbers(arr []int) int {
	count := 0

	for i := 0; i < len(arr); i++ {
		isUnique := true
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			count++
		}
	}

	return count
}
