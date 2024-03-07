package main

import (
	"fmt"
)

func main() {
	var T, P, solution, solutionT string

	fmt.Scanln(&P)
	fmt.Scanln(&T)
	printed := false

	for i := 0; i <= len(T)-len(P); i++ {
		if len(T) < len(P) || len(T) == 0 || len(P) == 0 {
			printed = true
			fmt.Println(0)
			break
		}
		for j := 0; j < len(P); j++ {
			solution += string(P[j])
			if P[j] == T[i+j] {
				solutionT += string(P[j])
			} else {
				break
			}
		}

		if P == solutionT {
			printed = true
			fmt.Println(solution)
			fmt.Println(i + 1)
			break
		} else {
			solutionT = ""
		}
	}
	if !printed {
		fmt.Println(solution)
		fmt.Println(0)
	}
}
