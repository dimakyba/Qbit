package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var arr []int = make([]int, 10)
	arr[0] = x
	arr[1] = y

	for i := 2; i < 10; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	for i := 0; i <= 4; i++ {
		Swap(arr, i, 9-i)
	}
	PrintArray(arr)
}

func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func Swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp

}
