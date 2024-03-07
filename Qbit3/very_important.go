package main

import (
	"fmt"
)

func main() {
	var n, k, i int64
	fmt.Scan(&n, &k)

	if n < k {
		fmt.Printf("%d very important numbers:\n", k-n+1)
		for i = n; i <= k; i++ {
			fmt.Println(i)
		}
	} else {
		fmt.Printf("%d very important numbers:\n", n-k+1)
		for i = n; i >= k; i-- {
			fmt.Println(i)
		}
	}
}
