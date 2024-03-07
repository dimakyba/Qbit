package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	andrew := make([]int, 0)
	pasha := make([]int, 0)
	wins1 := 0
	wins2 := 0

	for i := 0; i < n; i++ {
		var andrews, pashas int
		fmt.Scan(&andrews, &pashas)

		if andrews > pashas {
			wins1++
		} else if pashas > andrews {
			wins2++
		}

		andrew = append(andrew, andrews)
		pasha = append(pasha, pashas)
	}

	if wins1 > wins2 {
		fmt.Println("Andrey")
	} else if wins1 < wins2 {
		fmt.Println("Pasha")
	} else {
		sumAndrew := 0
		sumPasha := 0

		for i := 0; i < n; i++ {
			sumAndrew += andrew[i]
			sumPasha += pasha[i]
		}

		if sumAndrew > sumPasha {
			fmt.Println("Andrey")
		} else if sumPasha > sumAndrew {
			fmt.Println("Pasha")
		} else {
			fmt.Println("Draw")
		}
	}
}
