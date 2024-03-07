package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	modifiedString := ""

	for _, char := range s {
		switch char {
		case 'w':
			modifiedString += "z"
		case 'z':
			modifiedString += "w"
		default:
			modifiedString += string(char)
		}
	}

	fmt.Println(modifiedString)
}
