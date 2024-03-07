package main

import "fmt"

func main() {
	x, y := 0, 0
	dx, dy := 0, 1
	var c string

	fmt.Scan(&c)

	for c != "" {
		if c == "R" {
			temp := dx
			dx = dy
			dy = -temp
		} else if c == "L" {
			temp := dx
			dx = -dy
			dy = temp
		} else {
			x += dx
			y += dy
		}
		fmt.Scan(&c)
	}

	fmt.Printf("%d %d\n", x, y)
}
