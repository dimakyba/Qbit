package main
//really naive

import (
	"fmt"
)

func main() {
	var p, t string
	fmt.Scan(&p)
	fmt.Scan(&t)
	i := 0
	counter := 0
	pLength := len(p)
	tLength := len(t)

	for i <= tLength-pLength {
		counter++
		slice := t[i : i+pLength]
		if slice == p {
			fmt.Println(t[i:tLength])
			fmt.Println(counter)
			break
		}
		i++
	}
}
