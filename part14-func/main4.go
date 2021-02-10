package main

import "fmt"

func main() {
	defer first()
	second()
	second()
	second()
	second()
	second()
	second()
	second()
}

func first() {
	fmt.Println("0")
}
func second() {
	fmt.Println("1")
}
