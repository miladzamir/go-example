package main

import "fmt"

func main() {
	//5
	//5 * 4 * 3 * 2 * 1

	fmt.Println(foo(4))
}

func foo(number int) int {

	if number == 0 {
		return 1
	}
	return number * foo(number-1)
}
