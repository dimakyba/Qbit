package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println(2)
	} else if number%3 == 0 {
		fmt.Println(3)
	} else {
		is_printed := false
		for i := 5; i*i <= number; i += 2 {
			if number%i == 0 {
				fmt.Println(i)
				is_printed = true
				break
			}
		}
		if !is_printed {
			fmt.Println(number)
		}
	}
}
