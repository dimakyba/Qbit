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
	for i := 0; i < k; i++ {
		var n int
		fmt.Scan(&n)
		input := GetInputSlice()
		hasSameSign := HasTwoAdjacentElementsWithSameSign(input)

		if hasSameSign {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
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

func HasTwoAdjacentElementsWithSameSign(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if (arr[i] > 0 && arr[i+1] > 0) || (arr[i] < 0 && arr[i+1] < 0) {
			return true
		}
	}
	return false
}
