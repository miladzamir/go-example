package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mg = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	//bit shifting
	x := 2
	fmt.Printf("%d \t %b \n", x, x)

	y := x << 2
	fmt.Printf("%d \t %b \n", y, y)

	fmt.Println("-/-/-/-/-/-/-/-/-/-/-/-/-/-/-")

	fmt.Printf("%d \t\t %b \n", kb, kb)
	fmt.Printf("%d \t %b \n", mg, mg)
	fmt.Printf("%d \t %b \n", gb, gb)
}
