package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	counter := 0
	for n != 0 {
		counter++
		fmt.Scan(&n)
	}
	fmt.Println(counter)
}
