package main

import (
	"fmt"
)

func main() {
	var s, s1 string
	// s = "abracadabra"
	fmt.Scan(&s)
	var i, j int
	fmt.Scan(&i, &j)
	// i = 2
	// j = 5
	s1 = s[i:j]
	fmt.Println(j - i + 2)

	for i <= j {
		fmt.Println(s1)
		s1 = s1[1:]
		j--
	}
}
