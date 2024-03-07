package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	var k int
	fmt.Scan(&k)
	counter := 0
	length := len(s)

	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			counter++
		}
	}
	fmt.Println(counter)
}
