package main

import "fmt"

func main() {
	var s, s1 string
	fmt.Scan(&s)
	// s = "abracadabra"
	var i, j int
	fmt.Scan(&i, &j)
	// i = 1
	// j = 3
	i--
	j--
	fmt.Println(j - i + 2)
	for ; i < j; i++ {
		s1 += fmt.Sprintf("%c", s[i])
		fmt.Println(s1)
	}
}
