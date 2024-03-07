package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var s rune = ' '
	fmt.Scan(&N)

	solution := strings.Repeat(string(s), N)

	for i := 1; i <= N*2-1; i++ {
		if i <= N {
			solution = solution[:N-i] + "*" + solution[N-i+1:]
			if i >= 3 {
				solution = solution[:N-i+1] + "+" + solution[N-i+2:]
			}
			fmt.Println(solution)
		} else {
			solution = solution[:i-1-N] + " " + solution[i-N:]
			solution = solution[:i-N] + "*" + solution[i-N+1:]
			fmt.Println(solution)
		}
	}
}
