package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')

	var k int
	fmt.Scan(&k)
	counter := 0
	for i := 0; i < len(inputString); i++ {
		if inputString[i] == ' ' {
			counter++

			if counter == k {
				fmt.Println(i + 1)
				return
			}
		}
	}
	fmt.Println(0)
}

// doesnt work at all ;(
