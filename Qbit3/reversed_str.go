package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if s == "0" {
		fmt.Println(0)
	} else {

		s = strings.TrimRight(s, "0")
		if s[0] == '-' {
			s = strings.Trim(s, "-")
			s = reverse(s)
			s = "-" + s
		} else {
			s = reverse(s)
		}
		fmt.Println(s)

	}
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
