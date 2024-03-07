package main

import (
	"fmt"
	"math"
)

func main() {
	var price, uah, kop float64
	fmt.Scan(&price)
	price /= 10
	for i := 100; i <= 900; i += 100 {
		new_price := price * float64(i) / 100
		uah = math.Floor(new_price)
		kop = (new_price - uah) * 100
		fmt.Printf(" %d grams: %d UAH %d kop\n", i, int(math.Round(uah)), int(math.Round(kop)))
	}
	for i := 1000; i <= 1500; i += 100 {
		new_price := price * float64(i) / 100
		uah = math.Floor(new_price)
		kop = (new_price - uah) * 100
		fmt.Printf("%d grams: %d UAH %d kop\n", i, int(math.Round(uah)), int(math.Round(kop)))
	}
}
