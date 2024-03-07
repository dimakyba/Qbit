package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	fmt.Println("***")
	length := len(n)
	for i := 0; i < length; i++ {
		fmt.Printf("*%c*\n", n[i])
	}
	fmt.Println("***")
}
