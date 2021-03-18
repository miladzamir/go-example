package main

import "fmt"

func main() {
	// for {
	// 	fmt.Println("2021")
	// 	break
	// }
	year := 1980

	for {
		if year == 2021 {
			fmt.Println(year)
			break
		}
		year++
	}
}
