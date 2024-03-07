package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)
	var n, m int
	fmt.Scan(&n, &m)

	n--
	m--

	if n < 0 {
		n = 0
	}

	if m >= len(s) {
		m = len(s) - 1
	}

	resultString := s[:n] + s[m+1:]
	fmt.Println(resultString)
}
// doesn't work :(((
