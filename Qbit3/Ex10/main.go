package main

import "fmt"

func main() {
	var k, i int
	sum := 0
	fmt.Scan(&k)

	i = k
	sum += i
	for i != 0 || k != 0 {
		i = k
		fmt.Scan(&k)
		if i == 0 && k == 0 {
			break
		}
		sum += k
	}

	fmt.Println(sum)
}
