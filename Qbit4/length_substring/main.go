package main

import "fmt"

func main() {
	var s, s1 string
	// fmt.Scan(&s)
	s = "abracadabra"
	var i, j int
	// fmt.Scan(&i, &j)
	i = 5
	j = 3
	i--
	// j--
	if i < j {
		s1 = s[i:j]
		fmt.Println(len(s1))
		fmt.Println(s1)
	} else {
		fmt.Println(0)
	}

}
