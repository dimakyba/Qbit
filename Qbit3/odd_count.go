package main

import "fmt"

func main() {
	var n, digit int
	fmt.Scan(&n)
	counter := 0
	if n < 0 {
		n = -n
	}

	for n > 0 {
		digit = n % 10
		if digit%2 != 0 {
			counter++
		}
		n /= 10
	}

	fmt.Println(counter)
}
