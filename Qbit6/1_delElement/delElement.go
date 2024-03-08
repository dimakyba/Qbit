package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var idx int
	fmt.Scan(&idx)
	idx--

	doneArr := deleteElement(arr, idx)
	printArray(doneArr)
}

func printArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func deleteElement(arr []int, idx int) []int {
	arr2 := make([]int, len(arr)-1)
	for i := 0; i < idx; i++ {
		arr2[i] = arr[i]
	}
	for i := idx; i < len(arr)-1; i++ {
		arr2[i] = arr[i+1]
	}

	return arr2
}
