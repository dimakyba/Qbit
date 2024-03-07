package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Scan(&number)

	i := 0
	suma := 0.0
	var res1, res2 int

	for suma <= number+1e-6 {
		i++
		suma += 1.0 / float64(i)

		if suma < number-1e-6 {
			res1 = i
		}
	}

	res2 = i
	fmt.Println(res1, res2)
}
