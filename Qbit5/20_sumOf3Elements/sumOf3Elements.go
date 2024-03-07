package main

import (
	"fmt"
)

func main() {
	var x, p, q, r int
	fmt.Scan(&x)
	fmt.Scan(&p, &q, &r)

	arr := generateArray(x)
	sum := arr[p-1] + arr[q-1] + arr[r-1]
	fmt.Println(sum)

}

func generateArray(x int) []int {
	var arr []int = make([]int, 20)
	arr[0] = x

	for i := 2; i <= 20; i++ {
		arr[i-1] = arr[i-2]*(i-10) + x
	}
	return arr
}
