package main

import "fmt"

func main() {
	x := 0
	if x > 0 {
		fmt.Println("x > 0")
	} else if x > 0 {
		fmt.Println("x < 0")
	} else if x == 0 {
		fmt.Println("x = 0")
	} else {
		fmt.Println("x != 0")
	}
}
