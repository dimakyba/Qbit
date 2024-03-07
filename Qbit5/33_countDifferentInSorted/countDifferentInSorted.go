package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var size int
		fmt.Scan(&size)
		arr := GetInputSlice()
		fmt.Println(TwoNeighborsAreSameSign(arr))
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

func TwoNeighborsAreSameSign(arr []int) string {
	for i := 1; i < len(arr); i++ {
		if (arr[i-1] > 0 && arr[i] > 0) || (arr[i-1] < 0 && arr[i] < 0) {
			return "YES"
		}
	}
	return "NO"
}
