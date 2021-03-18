package main

import "fmt"

func main() {

	fmt.Println(foo())
	amount, _ := bar()
	fmt.Println(amount)
}

func foo() int {
	return 8585
}

func bar() (int, string) {
	return 100, "$"
}
