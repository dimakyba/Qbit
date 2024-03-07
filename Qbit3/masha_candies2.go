package main

import (
	"fmt"
)

func main() {
	var gryvnas, kopiyki int
	fmt.Scan(&gryvnas, &kopiyki)
	a := gryvnas
	b := kopiyki

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d kg: %d,%02d UAH\n", i, gryvnas, kopiyki)
		gryvnas += a
		kopiyki += b
		if kopiyki >= 100 {
			kopiyki %= 100
			gryvnas++
		}
	}
}
