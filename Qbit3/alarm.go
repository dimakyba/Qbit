package main

import (
	"fmt"
)

func main() {
	var counter, hours, minutes int
	fmt.Scan(&hours, &minutes)
	fmt.Scan(&counter)
	total_minutes := hours*60 + minutes

	for i := 1; i <= counter; i++ {
		if hours%24 != 0 || hours >= 24 {
			hours %= 24
		}
		if i == counter {
			fmt.Printf("%02d:%02d", hours, minutes)
		} else {
			fmt.Printf("%02d:%02d\n", hours, minutes)
		}

		total_minutes += 5
		hours = total_minutes / 60
		minutes = total_minutes % 60

	}
}
