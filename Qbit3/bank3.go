package main

import "fmt"

func main() {
	var currency, percent, finalCurrency float64
	fmt.Scan(&currency, &percent, &finalCurrency)

	year := 0
	cents := 0

	if currency <= 0 {
		fmt.Println(0)
	} else {
		for {
			if currency >= finalCurrency {
				fmt.Println(year)
				break
			}

			year++
			change := (currency + float64(cents)/100) / 100 * percent

			currency += float64(int(change))
			change -= float64(int(change))
			cents += int(change * 100)

			if cents >= 100 {
				currency += float64(cents / 100)
				cents %= 100
			}
		}
	}
}
