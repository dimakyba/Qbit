package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	time := make([][3]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&time[i][j])
		}
	}

	sort.Slice(time, func(i, j int) bool {
		return (time[i][0]*3600 + time[i][1]*60 + time[i][2]) < (time[j][0]*3600 + time[j][1]*60 + time[j][2])
	})

	for i := 0; i < n; i++ {
		h := time[i][0]
		m := time[i][1]
		s := time[i][2]

		fmt.Printf("%d %d %d\n", h, m, s)
	}
}
