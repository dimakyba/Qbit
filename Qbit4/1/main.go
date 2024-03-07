package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)
	nextChar := rune(s[0]) + 1
	previousChar := rune(s[0]) - 1
	fmt.Printf("The next char for '%c' (%d) is '%c' (%d).\n", s[0], s[0], nextChar, nextChar)
	fmt.Printf("The previous char for '%c' (%d) is '%c' (%d).\n", s[0], s[0], previousChar, previousChar)

}
