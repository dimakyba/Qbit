package main

import (
	"fmt"
)

func main() {
	var uah, kop float64
	fmt.Scan(&uah, &kop)
	total_price := uah + kop/100
	total_uah := int(total_price)
	total_kop := int((total_price - float64(total_uah)) * 100)
	fmt.Printf(" 1 kg: %d,%02d UAH\n", total_uah, total_kop)
	for i := 2; i <= 10; i++ {
		total_price = (uah + kop/100) * float64(i)
		total_uah = int(total_price)
		total_kop = int((total_price - float64(total_uah)) * 100)

		fmt.Printf("%d kg: %d,%02d UAH\n", i, total_uah, total_kop)
	}

}
