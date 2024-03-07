package main

import (
	"fmt"
)

func main() {
	var s, s1 string
	// fmt.Scan(&s)
	s = "abracadabra"
	var i, j int
	// fmt.Scan(&i, &j)
	i = 2
	j = 5
	// i--
	// j--

	s1 = s[i:j]
	length := len(s1)
	for i < j {
		fmt.Println(s1)
		s1 = s1[:length-1] // Removing the last character from s1
		j--                // Decrement j since the string length is decreasing
	}
}
