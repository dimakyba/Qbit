package main

import (
	"fmt"
	"math"
)

func main() {
	var n, p, m, i float64
	fmt.Scan(&n, &p, &m)
	p /= 100.0
	for i = 1; n*math.Pow(1+p, float64(i)) <= m; i++ {
	}
	fmt.Println(i)
}
