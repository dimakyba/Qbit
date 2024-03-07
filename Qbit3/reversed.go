package main

import (
	"fmt"
)

func main() {
	var num, reversedNum, digit int
	fmt.Scan(&num)

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	for num > 0 {
		digit = num % 10
		reversedNum = reversedNum*10 + digit
		num /= 10
	}

	if isNegative {
		reversedNum = -reversedNum
	}

	fmt.Println(reversedNum)
}
